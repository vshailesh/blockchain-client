package main

import (
	"os"
	"testing"
)

func TestGetPayload_Default(t *testing.T) {
	payload, err := getPayload("")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if payload.Method != "eth_blockNumber" {
		t.Errorf("Expected default method 'eth_blockNumber', got '%s'", payload.Method)
	}
}

func TestGetPayload_FileInput(t *testing.T) {
	// Create temporary test file
	tmpFile, err := os.CreateTemp("", "test-*.json")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpFile.Name())

	content := `{"jsonrpc":"2.0","method":"test_method","id":123}`
	if _, err := tmpFile.WriteString(content); err != nil {
		t.Fatal(err)
	}
	tmpFile.Close()

	payload, err := getPayload(tmpFile.Name())
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if payload.Method != "test_method" {
		t.Errorf("Expected method 'test_method', got '%s'", payload.Method)
	}
}

func TestGetPayload_StdinInput(t *testing.T) {
	// Mock stdin
	r, w, _ := os.Pipe()
	origStdin := os.Stdin
	os.Stdin = r
	defer func() { os.Stdin = origStdin }()

	go func() {
		defer w.Close()
		w.WriteString(`{"jsonrpc":"2.0","method":"stdin_method","id":456}`)
	}()

	payload, err := getPayload("")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if payload.Method != "stdin_method" {
		t.Errorf("Expected method 'stdin_method', got '%s'", payload.Method)
	}
}

func TestIsInputFromPipe(t *testing.T) {
	// Test piped input
	r, w, _ := os.Pipe()
	origStdin := os.Stdin
	os.Stdin = r
	defer func() { os.Stdin = origStdin }()

	w.WriteString("test")
	w.Close()

	if !isInputFromPipe() {
		t.Error("Expected true for piped input")
	}

	// Test terminal input
	os.Stdin = origStdin
	if isInputFromPipe() {
		t.Error("Expected false for terminal input")
	}
}