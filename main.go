package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "logs-go/docs" // import generated documentation
	"logs-go/internal/api/handlers"
	"logs-go/internal/config"
	"logs-go/internal/database"
)

// @title Logs Service API
// @version 1.0
// @description API для сервиса логов
// @BasePath /api
func main() {
	// Загрузка конфигурации
	cfg, err := config.Load()
	if err != nil {
		fmt.Printf("Error loading config: %v\n", err)
		os.Exit(1)
	}

	// Инициализация подключения к базе данных
	dbConfig, err := config.NewDatabaseConfig()
	if err != nil {
		fmt.Printf("Error loading database config: %v\n", err)
		os.Exit(1)
	}

	backgroundContext := context.Background()

	pool, err := database.NewConnection(backgroundContext, dbConfig.ConnectionString())
	if err != nil {
		fmt.Printf("Error connecting to database: %v\n", err)
		os.Exit(1)
	}
	defer pool.Close()

	// Запуск миграций
	if err := database.RunMigrations(dbConfig.ConnectionString(), "migrations"); err != nil {
		fmt.Printf("Error running migrations: %v\n", err)
		os.Exit(1)
	}

	// Создание маршрутизатора
	router := handlers.SetupRouter(cfg)

	// Создание HTTP сервера
	address := ":" + cfg.Port
	fmt.Printf("Starting server on %s\n", address)

	server := &http.Server{
		Addr:         address,
		Handler:      router,
		ReadTimeout:  time.Duration(cfg.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(cfg.WriteTimeout) * time.Second,
		IdleTimeout:  time.Duration(cfg.IdleTimeout) * time.Second,
	}

	// Канал для получения сигналов завершения
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	// Запуск сервера в горутине
	go func() {
		fmt.Printf("Server is ready to accept connections on %s\n", address)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Printf("Error starting server: %v\n", err)
			os.Exit(1)
		}
	}()

	// Ожидание сигнала завершения
	<-stop
	fmt.Println("Shutting down server...")

	// Создание контекста с таймаутом для graceful shutdown
	shutdownCtx, cancel := context.WithTimeout(context.Background(), time.Duration(cfg.ShutdownTimeout)*time.Second)
	defer cancel()

	// Graceful shutdown
	if err := server.Shutdown(shutdownCtx); err != nil {
		fmt.Printf("Error during server shutdown: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Server stopped")
}
