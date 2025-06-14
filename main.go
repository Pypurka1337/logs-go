package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"hello-world-api/internal/api/handlers"
	"hello-world-api/internal/config"
)

func main() {
	// Загрузка конфигурации
	cfg, err := config.Load()
	if err != nil {
		fmt.Printf("Error loading config: %v\n", err)
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
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(cfg.ShutdownTimeout)*time.Second)
	defer cancel()

	// Graceful shutdown
	if err := server.Shutdown(ctx); err != nil {
		fmt.Printf("Error during server shutdown: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Server stopped")
}
