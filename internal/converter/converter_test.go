package converter

import (
	"bytes"
	"encoding/json"
	"testing"
)

func TestCSVToJSON(t *testing.T) {
	input := `name,age
Alice,25
Bob,30`

	var buf bytes.Buffer
	err := CSVToJSON(bytes.NewReader([]byte(input)), &buf)
	if err != nil {
		t.Fatalf("CSVToJSON failed: %v", err)
	}

	// Декодируем результат в структуру
	var result []map[string]interface{}
	if err := json.Unmarshal(buf.Bytes(), &result); err != nil {
		t.Fatal(err)
	}

	// Проверяем содержимое
	if len(result) != 2 {
		t.Fatalf("Expected 2 records, got %d", len(result))
	}

	if result[0]["name"] != "Alice" || result[0]["age"] != 25.0 {
		t.Errorf("First record mismatch: %v", result[0])
	}

	if result[1]["name"] != "Bob" || result[1]["age"] != 30.0 {
		t.Errorf("Second record mismatch: %v", result[1])
	}
}
