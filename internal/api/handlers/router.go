package handlers

import (
	"encoding/json"
	"logs-go/internal/config"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/swaggo/http-swagger"
	_ "logs-go/docs" // import generated documentation
)

func SetupRouter(cfg *config.Config) *chi.Mux {
	router := chi.NewRouter()

	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	})
	router.Use(corsMiddleware.Handler)

	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.Timeout(time.Duration(cfg.WriteTimeout) * time.Second))

	router.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("/swagger/doc.json"),
	))

	router.Route("/api", func(r chi.Router) {
		r.Get("/health", healthHandler())
	})

	return router
}

// healthHandler godoc
// @Summary Проверка здоровья сервиса
// @Description Проверяет работоспособность сервиса
// @Tags health
// @Accept json
// @Produce json
// @Success 200 {object} map[string]string
// @Router /api/health [get]
func healthHandler() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusOK)
		json.NewEncoder(writer).Encode(map[string]string{
			"status": "OK",
		})
	}
}
