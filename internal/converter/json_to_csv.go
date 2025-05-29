package converter

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"sort"
)

// JSONToCSV конвертирует JSON в CSV потоково.
func JSONToCSV(jsonReader io.Reader, csvWriter io.Writer) error {
	decoder := json.NewDecoder(jsonReader)
	writer := csv.NewWriter(csvWriter)
	defer writer.Flush()

	// Читаем первый элемент, чтобы определить заголовки
	var firstRecord map[string]interface{}
	if err := decoder.Decode(&firstRecord); err != nil {
		return fmt.Errorf("JSON decode: %v", err)
	}

	// Получаем все возможные заголовки
	headers := make([]string, 0, len(firstRecord))
	for key := range firstRecord {
		headers = append(headers, key)
	}
	sort.Strings(headers) // Сортируем для стабильности

	// Пишем заголовки
	if err := writer.Write(headers); err != nil {
		return err
	}

	// Обрабатываем первую запись
	if err := writeCSVRecord(writer, headers, firstRecord); err != nil {
		return err
	}

	// Обрабатываем остальные записи
	for {
		var record map[string]interface{}
		if err := decoder.Decode(&record); err != nil {
			if err == io.EOF {
				break
			}
			return fmt.Errorf("JSON decode: %v", err)
		}

		if err := writeCSVRecord(writer, headers, record); err != nil {
			return err
		}
	}

	return nil
}

func writeCSVRecord(writer *csv.Writer, headers []string, record map[string]interface{}) error {
	row := make([]string, len(headers))
	for i, header := range headers {
		val := record[header]
		row[i] = convertValueToString(val)
	}
	return writer.Write(row)
}

func convertValueToString(v interface{}) string {
	if v == nil {
		return ""
	}

	switch val := v.(type) {
	case string:
		return val
	case bool, int, int8, int16, int32, int64, uint, uint8, uint32, uint64, float32, float64:
		return fmt.Sprintf("%v", val)
	case []interface{}, map[string]interface{}:
		jsonData, err := json.Marshal(val)
		if err != nil {
			return fmt.Sprintf("%v", val)
		}
		return string(jsonData)
	default:
		return fmt.Sprintf("%v", val)
	}
}
