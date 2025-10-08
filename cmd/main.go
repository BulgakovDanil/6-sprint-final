package main

import (
	"log"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

func main() {
	//Создаем переменную типа log.Logger
	var logger *log.Logger

	//Создаем сервер с помощью функции Router из пакета server
	server := server.Router(logger)

	//Запускаем сервер
	err := server.Server.ListenAndServe()
	if err != nil {
		//В случае ошибки запуска сервера логируем и завершаем приложение
		log.Fatal(err)
	}
}
