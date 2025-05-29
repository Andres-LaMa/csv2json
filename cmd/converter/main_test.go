package main

import (
	"flag"
	"os"
	"testing"
)

func TestMainLogic(t *testing.T) {
	// Сохраняем оригинальные аргументы
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()

	// Создаем тестовые файлы
	testCSV := "test.csv"
	testJSON := "test.json"
	os.WriteFile(testCSV, []byte("name,age\nAlice,25"), 0644)
	defer os.Remove(testCSV)
	defer os.Remove(testJSON)

	// Тестируем CSV -> JSON
	os.Args = []string{"cmd", "-input", testCSV, "-output", testJSON}
	flag.Parse()
	main()

	// Проверяем что выходной файл создан
	if _, err := os.Stat(testJSON); os.IsNotExist(err) {
		t.Error("Output file not created")
	}
}
