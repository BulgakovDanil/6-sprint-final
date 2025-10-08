package server

import (
	"log"
	"net/http"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/handlers"
)

// Servers структура содержащая конфигурацию http сервера и логгер
type Servers struct {
	Logger *log.Logger
	Server *http.Server
}

// Router функция возвращает экземпляр структуры для настройки сервера и регистрирует хендлеры
func Router(logger *log.Logger) *Servers {

	//Создаем роутер
	mux := http.NewServeMux()

	//Регистрируем хендлеры
	mux.HandleFunc("GET /", handlers.HTMLHandler)
	mux.HandleFunc("POST /upload", handlers.UploadHandler)

	//Создаем экземпляр структуры
	server := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ErrorLog:     logger,
		ReadTimeout:  time.Second * 5,
		WriteTimeout: time.Second * 10,
		IdleTimeout:  time.Second * 15,
	}

	//Возвращаем ссылку сервера
	return &Servers{
		Logger: logger,
		Server: server,
	}
}
