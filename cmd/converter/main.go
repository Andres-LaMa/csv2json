package main

import (
	"flag"
	"fmt"
	"os"

	"your_project/internal/converter"
	"your_project/internal/utils"
)

var (
	inputFile  = flag.String("input", "", "Input file (CSV or JSON)")
	outputFile = flag.String("output", "", "Output file")
	mode       = flag.String("mode", "csv2json", "Conversion mode: csv2json or json2csv")
	version    = "1.0.0" // Устанавливается при сборке через -ldflags
)

func main() {
	flag.Parse()
	utils.LogInfo(fmt.Sprintf("CSV↔JSON Converter v%s", version))

	if *inputFile == "" || *outputFile == "" {
		utils.LogError(fmt.Errorf("укажите --input и --output файлы"))
		os.Exit(1)
	}

	// Открываем входной файл
	input, err := os.Open(*inputFile)
	if err != nil {
		utils.LogError(err)
		os.Exit(1)
	}
	defer input.Close()

	// Создаем выходной файл
	output, err := os.Create(*outputFile)
	if err != nil {
		utils.LogError(err)
		os.Exit(1)
	}
	defer output.Close()

	// Выбираем режим конвертации
	switch *mode {
	case "csv2json":
		err = converter.CSVToJSON(input, output)
	case "json2csv":
		err = converter.JSONToCSV(input, output)
	default:
		err = fmt.Errorf("неподдерживаемый режим: %s", *mode)
	}

	if err != nil {
		utils.LogError(err)
		os.Exit(1)
	}

	utils.LogInfo("Конвертация завершена успешно!")
}
