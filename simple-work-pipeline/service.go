package main

// Upper, reverse, remove whitespace

import (
	"strings"
	"unicode"
)

type Service func(text string) string

func NewUpperService() Service {
	return func(text string) string {
		return strings.ToUpper(text)
	}
}

func NewReverseService() Service {
	return func(text string) string {
		buf := make([]byte, len(text))
		for i := len(text) - 1; i >= 0; i-- {
			buf[len(text)-1-i] = text[i]
		}
		return string(buf)
	}
}

func NewRemoveWhiteSpaceService() Service {
	return func(text string) string {
		var b strings.Builder
		b.Grow(len(text))
		for _, ch := range text {
			if !unicode.IsSpace(ch) {
				b.WriteRune(ch)
			}
		}
		return b.String()
	}
}
