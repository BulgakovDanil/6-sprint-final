package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

// HTMLHandler обрабатывает GET запрос к корневому пути "/" и возвращает HTML из файла index.html
func HTMLHandler(w http.ResponseWriter, r *http.Request) {
	//Получаем текущую директорию
	dir, _ := os.Getwd()

	var foundPath string

	for {
		path := filepath.Join(dir, "index.html")
		if _, err := os.Stat(path); err == nil {
			foundPath = path
			break
		}
	}

	//Читаем данные из файла "index.html"
	file, err := os.ReadFile(foundPath)
	if err != nil {
		//Если не удалось возвращаем ошибку
		http.Error(w, "file not found", http.StatusInternalServerError)
		return
	}

	//Устанвливаем заголовок
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	//Отправляем результат
	w.Write(file)
}

// UploadHandler обрабатывает POST запросы на пути "/upload" парсит загруженный файл,
// передает данные в функцию определения (TextDefinition) пакета service,
// сохраняет результат в локальный файл и возвращает его
func UploadHandler(w http.ResponseWriter, r *http.Request) {

	//Парсим файл с лимитом 10MB
	r.ParseMultipartForm(10 << 20)

	//Получаем файл из формы "myFile"
	file, header, err := r.FormFile("myFile")
	if err != nil {
		//Если не удалось возвращаем ошибку
		http.Error(w, "error when receiving the file", http.StatusInternalServerError)
		return
	}

	//Гарантируем отложенное закрытие файла
	defer file.Close()

	//Читаем все данные из загруженного файла в байтовый срез
	data, err := io.ReadAll(file)
	if err != nil {
		//Если не удалось возвращаем ошибку
		http.Error(w, "Error reading file", http.StatusInternalServerError)
		return
	}

	//Конвертируем байты в строку
	text := string(data)

	//Используем функцию TextDefinition из пакета service для определения текса
	result, err := service.TextDefinition(text)
	if err != nil {
		//Если не удалось возавращаем и логируем ошибку
		log.Printf("error in the definition: %v", err)
		return
	}

	//Получаем расширение файла
	ext := filepath.Ext(header.Filename)
	//Получаем текущее вреимя
	time := time.Now().UTC().Format("02-01-2006_15-04-05")
	//Генерируем имя файла на основе текущего времени и расширения файла пользователя
	nameFile := fmt.Sprintf("%s%s", time, ext)

	//Создаем локальный файл для сохранения результата
	localFile, err := os.Create(nameFile)
	if err != nil {
		//Если не удалось возвращаем ошибку
		http.Error(w, "Error creating file", http.StatusInternalServerError)
		return
	}

	//Записываем результат в локальный файл
	_, err = localFile.Write([]byte(result))
	if err != nil {
		//Если не удалось возвращаем ошибку
		http.Error(w, "error writing to file", http.StatusInternalServerError)
	}

	//Устанавливаем заголовок
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	//Отправляем результат конвертации
	w.Write([]byte(result))
}
