package internal

import (
	"os"
	"strings"
	"testing"
)

func BenchmarkTextHandler(b *testing.B) {
	for i := 0; i < b.N; i++ {
		f, _ := os.Open("../test/pasta.txt")
		defer f.Close()
		Uwuify(f, os.Stdout)
	}
}

func TestAddActionsModifiesString(t *testing.T) {
	input, inputCopy := "hello", "hello"

	addActions(&input, 1)
	if input == inputCopy {
		t.Errorf("%s == %s", input, inputCopy)
	}
}

func TestAddActionsAppendsCorrectString(t *testing.T) {
	input := "hello"

	addActions(&input, 1)
	if !contains(input, actions) {
		t.Errorf("%s does not contain action", input)
	}
}

func TestAddExclamationsModifiesString(t *testing.T) {
	input, inputCopy := "hello!", "hello!"

	addExclamations(&input, 1)
	if input == inputCopy {
		t.Errorf("%s == %s", input, inputCopy)
	}
}

func TestAddExclamationsAddsCorrectString(t *testing.T) {
	input := "hello!"

	addExclamations(&input, 1)
	if !contains(input, exclamations) {
		t.Errorf("%s does not contain action", input)
	}
}

func TestAddKaomojiModifiesString(t *testing.T) {
	input, inputCopy := "hello", "hello"

	addKaomoji(&input, 1)
	if input == inputCopy {
		t.Errorf("%s == %s", input, inputCopy)
	}
}

func TestAddKaomojiAppendsCorrectString(t *testing.T) {
	input := "hello"

	addKaomoji(&input, 1)
	if !contains(input, kaomoji) {
		t.Errorf("%s does not contain kaomoji", input)
	}
}

func TestReplaceText(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{
			input: "hello",
			want:  "hewwo",
		},
		{
			input: "herpes",
			want:  "hewpes",
		},
		{
			input: "HELLO",
			want:  "HEWWO",
		},
		{
			input: "HERPES",
			want:  "HEWPES",
		},
		{
			input: "nope",
			want:  "nyope",
		},
		{
			input: "Nope",
			want:  "Nyope",
		},
		{
			input: "NOPE",
			want:  "NYOPE",
		},
		{
			input: "Dove",
			want:  "Duv",
		},
		{
			input: "DOVE",
			want:  "DUV",
		},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			input := test.input
			replaceText(&input, 1)
			if input != test.want {
				t.Errorf("%s != %s", input, test.want)
			}
		})
	}

}

func TestAddStuttersModifiesString(t *testing.T) {
	input, inputCopy := "hello", "hello"

	addStutters(&input, 1)
	if input == inputCopy {
		t.Errorf("%s == %s", input, inputCopy)
	}
}

func TestAddStuttersAddsDashes(t *testing.T) {
	input := "hello"

	addStutters(&input, 1)

	if !strings.Contains(input, "-") {
		t.Errorf("%s contains no dashes", input)
	}
}

func contains(s string, arr []string) bool {
	for _, v := range arr {
		if strings.Contains(s, v) {
			return true
		}
	}
	return false
}
