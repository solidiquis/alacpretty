package ui

import (
	"fmt"
	"testing"
)

func TestFontShuffler(t *testing.T) {
	subject := fontShuffler()
	for _, i := range subject {
		fmt.Println(i)
	}
}
