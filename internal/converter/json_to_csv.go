package converter

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"time"

	"your_project/internal/parser"
	"your_project/internal/utils"
)

// JSONToCSV конвертирует JSON в CSV потоково.
func JSONToCSV(jsonReader io.Reader, csvWriter io.Writer) error {
	rows, err := parser.ParseJSON(jsonReader)
	if err != nil {
		return err
	}

	w := csv.NewWriter(csvWriter)
	defer w.Flush()

	// Записываем заголовки (первые ключи из первого объекта)
	var headers []string
	firstRow := <-rows
	if firstRow == nil {
		return nil
	}

	for key := range firstRow {
		headers = append(headers, key)
	}
	if err := w.Write(headers); err != nil {
		utils.LogError(err)
		return err
	}

	// Записываем первую строку
	if err := writeCSVRow(w, headers, firstRow); err != nil {
		return err
	}

	// Обрабатываем остальные строки
	for row := range rows {
		if err := writeCSVRow(w, headers, row); err != nil {
			utils.LogError(err)
			continue
		}
	}

	return nil
}

// writeCSVRow записывает одну строку в CSV.
func writeCSVRow(w *csv.Writer, headers []string, row map[string]interface{}) error {
	var record []string
	for _, key := range headers {
		val := row[key]
		record = append(record, toString(val))
	}
	return w.Write(record)
}

// toString преобразует значение в строку (с поддержкой дат, чисел и вложенных JSON).
func toString(v interface{}) string {
	switch v := v.(type) {
	case string:
		return v
	case int, float64:
		return fmt.Sprintf("%v", v)
	case time.Time:
		return v.Format("2006-01-02")
	default:
		// Вложенные объекты/массивы → JSON-строка
		jsonData, _ := json.Marshal(v)
		return string(jsonData)
	}
}
