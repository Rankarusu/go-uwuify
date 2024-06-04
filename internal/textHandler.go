package internal

import (
	"io"
	"math/rand"
	"regexp"
	"strings"
	"sync"
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

var kaomojiAscii = []string{
	";;w;;",
	"OwO",
	"UwU",
	">w<",
	"^w^",
	"^-^",
	":3",
	"x3",
	"xD",
	"XD",
}

var kaomojiUnicode = []string{
	"ÚwÚ",
	"(・`ω´・)",
	"(* ^ ω ^)",
	"(o^▽^o)",
	"(o･ω･o)",
	"(≧◡≦)",
	"(*´▽`*)",
	"(*≧ω≦*)",
	"o(≧▽≦)o",
	"(ꈍᴗꈍ)♡",
	"(〃￣ω￣〃)ゞ",
	"( ੭•͈ω•͈)੭",
	"(づ｡◕‿◕｡)づ",
	"≽^•⩊•^≼",
	"(⁄ ⁄>⁄ω⁄<⁄ ⁄)⁄",
	"૮(っ˶ᵔ  ᵕ  ᵔ˶)ა",
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

type Options struct {
	TextReplacements float64
	Stutters         float64
	Kaomoji          float64
	Exclamations     float64
	Actions          float64
	Unicode          bool
}

func Uwuify(r io.Reader, w io.Writer, o Options) error {
	content, err := io.ReadAll(r)
	if err != nil {
		return err
	}
	s := string(content)

	words := strings.Split(s, " ")
	wg := sync.WaitGroup{}

	for i := 0; i < len(words); i++ {
		wg.Add(1)
		go func(w *string, wg *sync.WaitGroup) {
			defer wg.Done()
			replaceText(w, o.TextReplacements)
			addStutters(w, o.Stutters)
			addKaomoji(w, o.Kaomoji, o.Unicode)
			addExclamations(w, o.Exclamations)
			addActions(w, o.Actions)
		}(&words[i], &wg)
	}
	wg.Wait()

	s = strings.Join(words, " ")

	_, err = w.Write([]byte(s))
	if err != nil {
		return err
	}

	return nil
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

func addKaomoji(s *string, chance float64, unicode bool) {
	pool := kaomojiAscii
	if unicode {
		pool = append(pool, kaomojiUnicode...)
	}
	if rand.Float64() < chance {
		*s = *s + " " + pool[rand.Intn(len(pool))]
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
