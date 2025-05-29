package parser

import (
	"encoding/csv"
	"io"
	"log"
	"strconv"
	"time"
)

// CSVRow представляет одну строку CSV (ключ-значение).
type CSVRow map[string]interface{}

// ParseCSV читает CSV потоково и возвращает канал строк.
func ParseCSV(r io.Reader) (<-chan CSVRow, error) {
	reader := csv.NewReader(r)
	rows := make(chan CSVRow)

	go func() {
		defer close(rows)

		// Читаем заголовки (если есть)
		headers, err := reader.Read()
		if err != nil {
			log.Printf("[ERROR] CSV headers: %v", err)
			return
		}

		for {
			record, err := reader.Read()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Printf("[ERROR] CSV read: %v", err)
				continue
			}

			row := make(CSVRow)
			for i, val := range record {
				if i >= len(headers) {
					continue
				}
				key := headers[i]
				row[key] = parseValue(val) // Автоматическое определение типа
			}
			rows <- row
		}
	}()

	return rows, nil
}

// parseValue преобразует строку в число, дату или оставляет строкой.
func parseValue(s string) interface{} {
	// Попытка распарсить как int
	if i, err := strconv.Atoi(s); err == nil {
		return i
	}

	// Попытка распарсить как float
	if f, err := strconv.ParseFloat(s, 64); err == nil {
		return f
	}

	// Попытка распарсить как дату (формат: 2006-01-02)
	if t, err := time.Parse("2006-01-02", s); err == nil {
		return t
	}

	// Возвращаем как строку
	return s
}
