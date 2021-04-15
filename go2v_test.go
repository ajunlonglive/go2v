package main

import (
	"os"
	"strings"
	"testing"

	"github.com/crthpl/go2v/convert"
)

func TestConvert(t *testing.T) {
	files, err := os.ReadDir("./tests")
	if err != nil {
		t.Error(err)
	}
	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".go") {
			converted, err := convert.Convert("./tests/" + file.Name())
			if err != nil {
				t.Fatal(err)
			}
			vfile, err := os.ReadFile("./tests/" + strings.TrimSuffix(file.Name(), ".go") + ".v")
			if err != nil {
				t.Fatal(err)
			}

			if string(vfile) != converted {
				t.Error("failed test, expected:\n" + string(vfile) + "\n----\ngot:\n" + converted)
			}
		}
	}
}
