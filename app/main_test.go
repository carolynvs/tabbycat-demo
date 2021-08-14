package main

import (
	"bytes"
	"fmt"
	"testing"
)

func TestList(t *testing.T) {
	w := &bytes.Buffer{}

	err := renderList(w)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(w.String())
}
