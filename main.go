package main

import (
	"errors"
	"io"

	"strings"
	"unicode"
)

var errUnexpectedCharacter = errors.New("unexpected character")

type Environment struct {
	name  string
	value string
}

func getString(r *[]rune) string {
	runes := *r
	var b strings.Builder
	for {
		if len(runes) <= 0 || !unicode.IsLetter(runes[0]) {
			break
		}

		b.WriteRune(runes[0])
		runes = runes[1:]
	}

	*r = runes

	return b.String()
}

func Parser(r io.Reader) (Environment, error) {
	b, err := io.ReadAll(r)
	if err != nil {
		return Environment{}, err
	}
	runes := []rune(string(b))

	var e Environment
	if !unicode.IsLetter(runes[0]) {
		return Environment{}, errUnexpectedCharacter
	}
	e.name = getString(&runes)

	if runes[0] != '=' {
		return Environment{}, errUnexpectedCharacter
	}
	runes = runes[1:]

	if !unicode.IsLetter(runes[0]) {
		return Environment{}, errUnexpectedCharacter
	}
	e.value = getString(&runes)

	return e, nil
}
