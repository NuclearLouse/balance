package logger

import (
	"balance/utilits/config"
	"bytes"
	"strings"
	"testing"
)

func TestLogger_New(t *testing.T) {
	cfg := config.New()
	cfg.Logger.Level = "error"
	log, err := New(cfg)
	if err != nil {
		t.Fatal(err)
	}
	buf := &bytes.Buffer{}
	log.Out = buf
	log.Error("error message")
	if !strings.Contains(buf.String(), "error message") {
		t.Fatal("expected log message")
	}
}
