package main

import (
	"io"
	"strings"
	"testing"
)

func TestParser(t *testing.T) {
	testData := struct {
		data     io.Reader
		expected Environment
	}{
		data: strings.NewReader("hoge=foo"),
		expected: Environment{
			name:  "hoge",
			value: "foo",
		},
	}
	e, err := Parser(testData.data)
	if err != nil {
		t.Error(err)
	}
	if e.name != testData.expected.name ||
		e.value != testData.expected.value {
		t.Fatalf("expected %###v, but got %###v\n", testData.expected, e)
	}
}
