package converter

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
)

// JSONToCSV конвертирует JSON в CSV потоково.
func JSONToCSV(jsonReader io.Reader, csvWriter io.Writer) error {
	decoder := json.NewDecoder(jsonReader)
	writer := csv.NewWriter(csvWriter)
	defer writer.Flush()

	// Проверяем первый токен, чтобы определить тип данных
	token, err := decoder.Token()
	if err != nil {
		return fmt.Errorf("JSON decode error: %v", err)
	}

	// Обрабатываем разные случаи ввода
	switch tok := token.(type) {
	case json.Delim:
		if tok == '[' {
			// Это массив - обрабатываем каждый элемент
			return processJSONArray(decoder, writer)
		} else if tok == '{' {
			// Это объект - обрабатываем как единичную запись
			return processJSONObject(decoder, writer, tok)
		}
	default:
		return fmt.Errorf("unexpected JSON token: %v", tok)
	}

	return nil
}

func processJSONArray(decoder *json.Decoder, writer *csv.Writer) error {
	var headers []string
	firstItem := true

	for decoder.More() {
		var record map[string]interface{}
		if err := decoder.Decode(&record); err != nil {
			return fmt.Errorf("JSON decode error: %v", err)
		}

		if firstItem {
			// Определяем заголовки по первой записи
			headers = make([]string, 0, len(record))
			for key := range record {
				headers = append(headers, key)
			}
			if err := writer.Write(headers); err != nil {
				return err
			}
			firstItem = false
		}

		if err := writeCSVRow(writer, headers, record); err != nil {
			return err
		}
	}

	// Прочитать закрывающую скобку массива
	_, err := decoder.Token()
	return err
}

func processJSONObject(decoder *json.Decoder, writer *csv.Writer, startToken json.Delim) error {
	var record map[string]interface{}
	if err := decoder.Decode(&record); err != nil {
		return fmt.Errorf("JSON decode error: %v", err)
	}

	// Определяем заголовки
	headers := make([]string, 0, len(record))
	for key := range record {
		headers = append(headers, key)
	}

	// Пишем заголовки и данные
	if err := writer.Write(headers); err != nil {
		return err
	}
	if err := writeCSVRow(writer, headers, record); err != nil {
		return err
	}

	// Прочитать закрывающую скобку объекта
	_, err := decoder.Token()
	return err
}

func writeCSVRow(writer *csv.Writer, headers []string, record map[string]interface{}) error {
	row := make([]string, len(headers))
	for i, header := range headers {
		val := record[header]
		switch v := val.(type) {
		case map[string]interface{}, []interface{}:
			jsonData, err := json.Marshal(v)
			if err != nil {
				return err
			}
			row[i] = string(jsonData)
		default:
			row[i] = fmt.Sprintf("%s", val)
		}
	}
	return writer.Write(row)
}
