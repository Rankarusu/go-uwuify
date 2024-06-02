package internal

import (
	"fmt"
	"os"
	"strings"
	"testing"
)

var probabilities = []struct {
	probability   float64
	shouldBeEqual bool
}{
	{0, true},
	{1, false},
}

func BenchmarkTextHandler(b *testing.B) {
	for i := 0; i < b.N; i++ {
		f, _ := os.Open("../test/pasta.txt")
		defer f.Close()
		Uwuify(f, os.Stdout, Options{1, 1, 1, 1, 1, true})
	}
}

func TestAddActionsModifiesString(t *testing.T) {

	for _, test := range probabilities {
		t.Run(fmt.Sprintf("%f", test.probability), func(t *testing.T) {
			input, inputCopy := "hello", "hello"

			addActions(&input, test.probability)
			if res := input == inputCopy; res != test.shouldBeEqual {
				t.Errorf("%s, %s, should be equal: %t", input, inputCopy, test.shouldBeEqual)
			}
		})
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

	for _, test := range probabilities {
		t.Run(fmt.Sprintf("%f", test.probability), func(t *testing.T) {
			input, inputCopy := "hello!", "hello!"

			addExclamations(&input, test.probability)
			if res := input == inputCopy; res != test.shouldBeEqual {
				t.Errorf("%s, %s, should be equal: %t", input, inputCopy, test.shouldBeEqual)
			}
		})
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
	for _, test := range probabilities {
		t.Run(fmt.Sprintf("%f", test.probability), func(t *testing.T) {
			input, inputCopy := "hello!", "hello!"

			addKaomoji(&input, test.probability, true)
			if res := input == inputCopy; res != test.shouldBeEqual {
				t.Errorf("%s, %s, should be equal: %t", input, inputCopy, test.shouldBeEqual)
			}
		})
	}
}

func TestAddKaomojiAppendsCorrectString(t *testing.T) {
	input := "hello"

	addKaomoji(&input, 1, false)

	if contains(input, kaomojiUnicode) {
		t.Errorf("%s contains unicode kaomoji", input)
	}

	if !contains(input, kaomojiAscii) {
		t.Errorf("%s does not contain kaomoji", input)
	}
}

func TestReplaceText(t *testing.T) {
	tests := []struct {
		input       string
		probability float64
		want        string
	}{
		{
			input:       "hello",
			probability: 1,
			want:        "hewwo",
		},
		{
			input:       "herpes",
			probability: 1,
			want:        "hewpes",
		},
		{
			input:       "HELLO",
			probability: 1,
			want:        "HEWWO",
		},
		{
			input:       "HERPES",
			probability: 1,
			want:        "HEWPES",
		},
		{
			input:       "nope",
			probability: 1,
			want:        "nyope",
		},
		{
			input:       "Nope",
			probability: 1,
			want:        "Nyope",
		},
		{
			input:       "NOPE",
			probability: 1,
			want:        "NYOPE",
		},
		{
			input:       "Dove",
			probability: 1,
			want:        "Duv",
		},
		{
			input:       "DOVE",
			probability: 1,
			want:        "DUV",
		},
		{
			input:       "hello",
			probability: 0,
			want:        "hello",
		},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			input := test.input
			replaceText(&input, test.probability)
			if input != test.want {
				t.Errorf("%s != %s", input, test.want)
			}
		})
	}
}

func TestAddStuttersModifiesString(t *testing.T) {
	for _, test := range probabilities {
		t.Run(fmt.Sprintf("%f", test.probability), func(t *testing.T) {
			input, inputCopy := "hello!", "hello!"

			addStutters(&input, test.probability)
			if res := input == inputCopy; res != test.shouldBeEqual {
				t.Errorf("%s, %s, should be equal: %t", input, inputCopy, test.shouldBeEqual)
			}
		})
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
