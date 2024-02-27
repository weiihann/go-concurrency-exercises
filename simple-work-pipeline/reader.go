package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

// A Reader should read and output lines
type Reader interface {
	Read(filepath string)
	Next() (string, bool)
	Close() error
}

func SelectReader(filePath string) Reader {
	if strings.Contains(filePath, ".txt") {
		return NewTextReader()
	} else {
		return nil
	}
}

type TextReader struct {
	scanner *bufio.Scanner
	file    *os.File
}

func NewTextReader() *TextReader {
	return &TextReader{}
}

func (t *TextReader) Read(filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("create new text reader failed, err: %s\n", err)
		return
	}

	t.scanner = bufio.NewScanner(file)
	t.file = file
}

func (t *TextReader) Next() (string, bool) {
	exist := t.scanner.Scan()
	if !exist {
		t.Close()
		return "", false
	}
	return t.scanner.Text(), true
}

func (t *TextReader) Close() error {
	return t.file.Close()
}
