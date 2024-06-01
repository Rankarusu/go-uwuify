package internal

import (
	"io"
	"log"
	"math/rand"
	"regexp"
	"strings"
	"unicode"
)

var exclamationsPattern = regexp.MustCompile(`[.!?]+`)

var exclamations = []string{"!?", "?!!", "?!?1", "!!11", "?!?!"}

var actions = []string{
	"*blushes*",
	"*whispers to self*",
	"*cries*",
	"*screams*",
	"*sweats*",
	"*twerks*",
	"*runs away*",
	"*screeches*",
	"*walks away*",
	"*sees bulge*",
	"*looks at you*",
	"*notices bulge*",
	"*starts twerking*",
	"*huggles tightly*",
	"*boops your nose*",
}

var kaomoji = []string{
	"(・`ω´・)",
	";;w;;",
	"OwO",
	"UwU",
	">w<",
	"^w^",
	"ÚwÚ",
	"^-^",
	":3",
	"x3",
}

var textreplacementMap = []struct {
	pattern      *regexp.Regexp
	replaceValue string
}{
	{regexp.MustCompile("(?:[rl])"), "w"},
	{regexp.MustCompile("(?:[RL])"), "W"},
	{regexp.MustCompile("n([aeiou])"), "ny$1"},
	{regexp.MustCompile("N([aeiou])"), "Ny$1"},
	{regexp.MustCompile("N([AEIOU])"), "NY$1"},
	{regexp.MustCompile("ove"), "uv"},
	{regexp.MustCompile("OVE"), "UV"},
}

func Uwuify(r io.Reader, w io.Writer) {
	content, err := io.ReadAll(r)
	if err != nil {
		log.Fatalf("error reading file - %s", err)
	}
	s := string(content)

	words := strings.Split(s, " ")
	for i := 0; i < len(words); i++ {
		replaceText(&words[i], .5)
		addStutters(&words[i], .025)
		addKaomoji(&words[i], .025)
		addExclamations(&words[i], .5)
		addActions(&words[i], .025)
	}

	s = strings.Join(words, " ")

	_, err = w.Write([]byte(s))
	if err != nil {
		log.Fatalf("error writing to file or stdout - %s", err)
	}
}

func addActions(s *string, chance float64) {
	if rand.Float64() < chance {
		*s = *s + " " + actions[rand.Intn(len(actions))]
	}
}

func addExclamations(s *string, chance float64) {
	*s = exclamationsPattern.ReplaceAllStringFunc(*s, func(s string) string {
		if rand.Float64() < chance {
			return exclamations[rand.Intn(len(exclamations))]
		}
		return s
	})
}

func addKaomoji(s *string, chance float64) {
	if rand.Float64() < chance {
		*s = *s + " " + kaomoji[rand.Intn(len(kaomoji))]
	}
}

func replaceText(s *string, chance float64) {
	for _, v := range textreplacementMap {
		*s = v.pattern.ReplaceAllStringFunc(*s, func(s string) string {
			if rand.Float64() < chance {
				submatches := v.pattern.FindStringSubmatchIndex(s)
				replacement := []byte{}
				replacement = v.pattern.ExpandString(replacement, v.replaceValue, s, submatches)
				return string(replacement)
			}
			return s
		})

	}
}

func addStutters(s *string, chance float64) {
	firstLetter := (*s)[0]
	if unicode.IsLetter(rune(firstLetter)) && rand.Float64() < chance {
		stutter := rand.Intn(2) + 1
		*s = strings.Repeat(string(firstLetter)+"-", stutter) + *s
	}
}
