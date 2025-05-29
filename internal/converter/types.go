package converter

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"
)

// CSVRecord представляет одну запись CSV с поддержкой разных типов данных
type CSVRecord map[string]interface{}

// JSONRecord представляет одну запись JSON для конвертации
type JSONRecord map[string]interface{}

// InputFormat определяет поддерживаемые форматы ввода
type InputFormat string

const (
	FormatCSV  InputFormat = "csv"
	FormatJSON InputFormat = "json"
)

// Config содержит параметры конвертера
type Config struct {
	InputPath  string
	OutputPath string
	Format     InputFormat
	Delimiter  rune
	Pretty     bool // Форматированный вывод JSON
}

// ParseValue преобразует строку в соответствующий тип данных
func ParseValue(s string) interface{} {
	// Попытка распарсить как целое число
	if i, err := strconv.Atoi(s); err == nil {
		return i
	}

	// Попытка распарсить как число с плавающей точкой
	if f, err := strconv.ParseFloat(s, 64); err == nil {
		return f
	}

	// Попытка распарсить как булево значение
	if b, err := strconv.ParseBool(s); err == nil {
		return b
	}

	// Попытка распарсить как дату (несколько форматов)
	dateFormats := []string{
		time.RFC3339,
		"2006-01-02",
		"02.01.2006",
	}
	for _, format := range dateFormats {
		if t, err := time.Parse(format, s); err == nil {
			return t
		}
	}

	// Попытка распарсить как вложенный JSON
	var js json.RawMessage
	if err := json.Unmarshal([]byte(s), &js); err == nil {
		var result interface{}
		if err := json.Unmarshal(js, &result); err == nil {
			return result
		}
	}

	// Возвращаем как строку по умолчанию
	return s
}

// StringToInterface преобразует строку CSV в значение с автоматическим определением типа
func StringToInterface(s string) interface{} {
	return ParseValue(s)
}

// InterfaceToString преобразует значение обратно в строку для CSV
func InterfaceToString(v interface{}) string {
	switch val := v.(type) {
	case string:
		return val
	case int, int32, int64, float32, float64:
		return fmt.Sprintf("%v", val)
	case bool:
		return strconv.FormatBool(val)
	case time.Time:
		return val.Format(time.RFC3339)
	case nil:
		return ""
	default:
		// Для сложных типов (массивы, объекты) - сериализуем в JSON
		if b, err := json.Marshal(val); err == nil {
			return string(b)
		}
		return fmt.Sprintf("%v", val)
	}
}
