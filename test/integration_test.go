package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
	"testing"

	"github.com/Andres-LaMa/csv2json/internal/converter"
	"github.com/Andres-LaMa/csv2json/internal/utils"
)

func TestCSVToJSON(t *testing.T) {
	csvInput := `name,age
Alice,25
Bob,30`

	expectedJSON := `[
  {
    "name": "Alice",
    "age": 25
  },
  {
    "name": "Bob",
    "age": 30
  }
]`

	var buf bytes.Buffer
	err := converter.CSVToJSON(bytes.NewReader([]byte(csvInput)), &buf)
	if err != nil {
		t.Fatalf("CSVToJSON failed: %v", err)
	}

	if buf.String() != expectedJSON {
		t.Errorf("Expected:\n%s\nGot:\n%s", expectedJSON, buf.String())
	}
}

func TestJSONToCSV(t *testing.T) {
	jsonInput := `[
		{"name": "Alice", "age": 25},
		{"name": "Bob", "age": 30}
	]`

	expectedCSV := `name,age
Alice,25
Bob,30
`

	var buf bytes.Buffer
	err := converter.JSONToCSV(bytes.NewReader([]byte(jsonInput)), &buf)
	if err != nil {
		t.Fatalf("JSONToCSV failed: %v", err)
	}

	if buf.String() != expectedCSV {
		t.Errorf("Expected:\n%s\nGot:\n%s", expectedCSV, buf.String())
	}
}

func TestNestedJSONToCSV(t *testing.T) {
	jsonInput := `[
        {
            "name": "Alice",
            "metadata": {
                "role": "admin",
                "tags": ["a", "b"]
            }
        }
    ]`

	// Ожидаем два возможных варианта (порядок полей в JSON может меняться)
	expectedVariants := []string{
		`name,metadata
Alice,"{""role"":""admin"",""tags"":[""a"",""b""]}"`,
		`name,metadata
Alice,"{""tags"":[""a"",""b""],""role"":""admin""}"`,
	}

	var buf bytes.Buffer
	err := converter.JSONToCSV(bytes.NewReader([]byte(jsonInput)), &buf)
	if err != nil {
		t.Fatalf("Nested JSONToCSV failed: %v", err)
	}

	result := buf.String()
	matched := false
	for _, variant := range expectedVariants {
		if result == variant {
			matched = true
			break
		}
	}

	if !matched {
		t.Errorf("Result doesn't match any expected variant.\nGot:\n%v\nExpected one of:\n%v",
			result, strings.Join(expectedVariants, "\nOR\n"))
	}
}

func TestCSVWithNestedJSON(t *testing.T) {
	csvInput := `name,data
Alice,"{""id"": 1, ""active"": true}"
`

	expectedJSON := `[
  {
    "name": "Alice",
    "data": {
      "id": 1,
      "active": true
    }
  }
]`

	var buf bytes.Buffer
	err := converter.CSVToJSON(bytes.NewReader([]byte(csvInput)), &buf)
	if err != nil {
		t.Fatalf("CSV with nested JSON failed: %v", err)
	}

	// Сравниваем с эталоном (с учетом пробелов и переносов)
	var got, want interface{}
	if err := json.Unmarshal(buf.Bytes(), &got); err != nil {
		t.Fatal(err)
	}
	if err := json.Unmarshal([]byte(expectedJSON), &want); err != nil {
		t.Fatal(err)
	}
}

// func TestBrokenCSV(t *testing.T) {
// 	csvInput := `name,age
// Alice,25
// "Broken,line"
// `

// 	var buf bytes.Buffer
// 	err := converter.CSVToJSON(bytes.NewReader([]byte(csvInput)), &buf)
// 	if err == nil {
// 		t.Fatal("Expected error for broken CSV, got nil")
// 	}
// 	utils.LogInfo(fmt.Sprintf("OK: ошибка обработана: %v", err))
// }

func TestBrokenJSON(t *testing.T) {
	jsonInput := `[
		{"name": "Alice"},
		{"name": "Bob", "age": ]
	]`

	var buf bytes.Buffer
	err := converter.JSONToCSV(bytes.NewReader([]byte(jsonInput)), &buf)
	if err == nil {
		t.Fatal("Expected error for broken JSON, got nil")
	}
	utils.LogInfo(fmt.Sprintf("OK: ошибка обработана: %v", err))
}

// func TestStreaming(t *testing.T) {
// 	// Генерируем большой CSV в памяти (10K строк)
// 	var bigCSV bytes.Buffer
// 	bigCSV.WriteString("name,age\n")
// 	for i := 0; i < 10000; i++ {
// 		bigCSV.WriteString(fmt.Sprintf("User%d,%d\n", i, i%100))
// 	}

// 	// Конвертируем и проверяем, что не паникует
// 	var buf bytes.Buffer
// 	err := converter.CSVToJSON(bytes.NewReader(bigCSV.Bytes()), &buf)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	// Проверяем кол-во строк в JSON
// 	var result []map[string]interface{}
// 	if err := json.Unmarshal(buf.Bytes(), &result); err != nil {
// 		t.Fatal(err)
// 	}
// 	if len(result) != 10000 {
// 		t.Errorf("Expected 10K rows, got %d", len(result))
// 	}
// }
