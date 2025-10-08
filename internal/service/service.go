package service

import (
	"fmt"
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

// TextDefinition функция автоматического определения кода Морза или обычного текста из переданной строки
func TextDefinition(text string) (string, error) {

	//Проверяем файл на пустоту
	if len(text) == 0 {
		//Если файл пустой возвращаем ошибку
		return "", fmt.Errorf("an empty file")
	}

	//Создаем переменную для хранения результата
	var result string

	//Удаляем левый символ
	cleaned := strings.TrimLeft(text, ".-")

	//Сравниваем длину очиещенного файла и исходного
	if len(cleaned) != len(text) {
		//Если длина очищенного файла не равна исходному конвертируем в код Морзе
		result = morse.ToText(text)
	} else {
		//Иначе конвертируем в Текст
		result = morse.ToMorse(text)
	}

	//Возвращаем результат и nil
	return result, nil

}
