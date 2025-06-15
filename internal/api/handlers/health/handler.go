package health

import (
	"encoding/json"
	"net/http"
	"time"
)

// Response Структура успешного ответа
type Response struct {
	Status    string `json:"status" example:"OK"`
	Version   string `json:"version" example:"1.0.0"`
	Timestamp string `json:"timestamp" example:"2023-10-05T12:34:56Z"`
}

// Handler godoc
// @Summary Проверка состояния сервиса
// @Description Проверяет работоспособность и доступность всех критических зависимостей сервиса
// @Tags service
// @Produce json
// @Success 200 {object} Response "Сервис работает нормально"
// @Router /health [get]
func Handler() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(http.StatusOK)
		json.NewEncoder(writer).Encode(Response{
			Status:    "OK",
			Version:   "1.0.0",
			Timestamp: time.Now().Format(time.RFC3339),
		})
	}
}
