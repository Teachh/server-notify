package logger_test

import (
	"bytes"
	"log"
	"testing"

	logger "github.com/Teachh/server-notify/internal/logger"
)

func TestLogger(t *testing.T) {
	var buf bytes.Buffer

	logger.Info = log.New(&buf, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	logger.Warning = log.New(&buf, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	logger.Error = log.New(&buf, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)

	logger.Info.Println("This is an info message")
	logger.Warning.Println("This is a warning message")
	logger.Error.Println("This is an error message")

	logOutput := buf.String()

	if !bytes.Contains(buf.Bytes(), []byte("INFO: ")) || !bytes.Contains(buf.Bytes(), []byte("This is an info message")) {
		t.Errorf("Info log not written correctly. Got: %s", logOutput)
	}
	if !bytes.Contains(buf.Bytes(), []byte("WARNING: ")) || !bytes.Contains(buf.Bytes(), []byte("This is a warning message")) {
		t.Errorf("Warning log not written correctly. Got: %s", logOutput)
	}
	if !bytes.Contains(buf.Bytes(), []byte("ERROR: ")) || !bytes.Contains(buf.Bytes(), []byte("This is an error message")) {
		t.Errorf("Error log not written correctly. Got: %s", logOutput)
	}
}
