package utils

import (
	"bytes"
	"testing"
)

func TestLoggers(t *testing.T) {
	// Сохраняем оригинальные настройки логгеров
	originalInfoOutput := InfoLogger.Writer()
	originalErrorOutput := ErrorLogger.Writer()
	defer func() {
		InfoLogger.SetOutput(originalInfoOutput)
		ErrorLogger.SetOutput(originalErrorOutput)
	}()

	// Тестируем Info лог
	t.Run("InfoLogger", func(t *testing.T) {
		var buf bytes.Buffer
		InfoLogger.SetOutput(&buf)

		LogInfo("test info")
		// if !bytes.Contains(buf.Bytes(), []byte("[INFO] test info")) {
		// 	t.Error("Info log not working")
		// }
	})

	// Тестируем Error лог
	t.Run("ErrorLogger", func(t *testing.T) {
		var buf bytes.Buffer
		ErrorLogger.SetOutput(&buf)

		LogError("test error")
		// if !bytes.Contains(buf.Bytes(), []byte("[ERROR] test error")) {
		// 	t.Error("Error log not working")
		// }
	})
}
