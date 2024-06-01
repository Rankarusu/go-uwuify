package cmd

import (
	"io"
	"log"
	"math/rand"
	"os"
	"regexp"
	"strings"
	"unicode"

	"github.com/spf13/cobra"
)

var spacePattern = regexp.MustCompile(`(\s)|($)`)
var exclamationsPattern = regexp.MustCompile(`[.!?]+`)
var wordStartPattern = regexp.MustCompile(`(^\w)|\s(\w)`)

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
	{regexp.MustCompile("N([AEIOU])"), "Ny$1"},
	{regexp.MustCompile("ove"), "uv"},
}

var rootCmd = &cobra.Command{
	Use:   "uwuify",
	Short: "uwuifies the given text",
	Long: `uwuifies the given text or file with options,
	e.g. uwuify -i input.txt -o output.txt`,
	Run: func(cmd *cobra.Command, args []string) {
		text, _ := cmd.Flags().GetString("text")
		infile, _ := cmd.Flags().GetString("infile")
		outfile, _ := cmd.Flags().GetString("outfile")

		var reader io.Reader

		stat, _ := os.Stdin.Stat()

		switch {
		case stat.Mode()&os.ModeDevice == 0:
			// if the input does not come from a character device (e.g. a terminal), it is most likely piped in
			reader = cmd.InOrStdin()
		case infile != "":
			file, err := os.Open(infile)
			if err != nil {
				log.Fatalf("error opening file - %s", err)
			}
			defer file.Close()
			reader = file
		case text != "":
			reader = strings.NewReader(text)
		default:
			cmd.Help()
			return
		}

		var writer io.Writer

		if outfile != "" {
			file, err := os.Create(outfile)
			if err != nil {
				log.Fatalf("error creating file - %s", err)
			}
			defer file.Close()
			writer = file
		} else {
			writer = cmd.OutOrStdout()
		}

		uwuify(reader, writer)
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringP("text", "t", "", "a text to uwuify")
	rootCmd.PersistentFlags().StringP("infile", "i", "", "a file to uwuify")
	rootCmd.PersistentFlags().StringP("outfile", "o", "", "a file to output the uwuified text to")
	rootCmd.Aliases = append(rootCmd.Aliases, "uwu")
}

func uwuify(r io.Reader, w io.Writer) {
	content, err := io.ReadAll(r)
	if err != nil {
		log.Fatalf("error reading file - %s", err)
	}
	s := string(content)

	replaceText(&s, .5)
	addStutters(&s, .025)
	addKaomoji(&s, .025)
	addExclamations(&s, .5)
	addActions(&s, .025)

	_, err = w.Write([]byte(s))
	if err != nil {
		log.Fatalf("error writing to file or stdout - %s", err)
	}
}

func addActions(s *string, chance float64) {
	*s = spacePattern.ReplaceAllStringFunc(*s, func(s string) string {
		if rand.Float64() < chance {
			return " " + actions[rand.Intn(len(actions))] + " "
		}
		return s
	})
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
	*s = spacePattern.ReplaceAllStringFunc(*s, func(s string) string {
		if rand.Float64() < chance {
			return " " + kaomoji[rand.Intn(len(kaomoji))] + " "
		}
		return s
	})
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
	split := strings.Split(*s, " ")
	for i := 0; i < len(split); i++ {
		if rand.Float64() < chance && unicode.IsLetter(rune(split[i][0])) {
			stutter := rand.Intn(2) + 1
			split[i] = strings.Repeat(string(split[i][0])+"-", stutter) + split[i]
		}
	}
	*s = strings.Join(split, " ")
}
