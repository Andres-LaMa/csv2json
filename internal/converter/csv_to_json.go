package converter

import (
	"encoding/json"
	"io"

	"github.com/Andres-LaMa/csv2json/internal/parser"
)

// CSVToJSON конвертирует CSV в JSON потоково.
func CSVToJSON(csvReader io.Reader, jsonWriter io.Writer) error {
	rows, err := parser.ParseCSV(csvReader)
	if err != nil {
		return err
	}

	jsonWriter.Write([]byte("[\n")) // Начало массива JSON
	first := true

	for row := range rows {
		if !first {
			jsonWriter.Write([]byte(",\n"))
		}
		first = false

		jsonData, err := json.MarshalIndent(row, "", "  ")
		if err != nil {
			return err
		}
		jsonWriter.Write(jsonData)
	}

	jsonWriter.Write([]byte("\n]")) // Конец массива JSON
	return nil
}
