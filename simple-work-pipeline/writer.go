package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

type Writer interface {
	Write(text string) error
	Close() error
}

type TextWriter struct {
	writer *bufio.Writer
	file   *os.File
}

func SelectWriter(filePath string) Writer {
	if strings.Contains(filePath, ".txt") {
		return NewTextWriter(filePath)
	} else {
		log.Fatalf("file format not supported: %s\n", filePath)
		return nil
	}
}

func NewTextWriter(filePath string) *TextWriter {
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		log.Fatalf("failed opening file: %s\n", err)
		return nil
	}

	bufWriter := bufio.NewWriter(file)

	return &TextWriter{
		writer: bufWriter,
		file:   file,
	}
}

func (tw *TextWriter) Write(text string) error {
	_, err := tw.writer.WriteString(text + "\n")
	return err
}

func (tw *TextWriter) Close() error {
	tw.writer.Flush()
	return tw.file.Close()
}
