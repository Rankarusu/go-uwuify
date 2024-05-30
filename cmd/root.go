package cmd

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"regexp"

	"github.com/spf13/cobra"
)

var spacePattern = regexp.MustCompile(" ")

var exclamationsPattern = regexp.MustCompile("[.!?]+")
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
	{
		pattern:      regexp.MustCompile("(?:[rl])"),
		replaceValue: "w",
	},
	{
		pattern:      regexp.MustCompile("(?:[RL])"),
		replaceValue: "W",
	},
	{
		pattern:      regexp.MustCompile("n([aeiou])"),
		replaceValue: "ny$1",
	},
	{
		pattern:      regexp.MustCompile("N([aeiou])"),
		replaceValue: "Ny$1",
	},
	{
		pattern:      regexp.MustCompile("N([AEIOU])"),
		replaceValue: "Ny$1",
	},
	{
		pattern:      regexp.MustCompile("ove"),
		replaceValue: "uv",
	},
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

		if text == "" && infile == "" && outfile == "" {
			cmd.Help()
		}

		var res string
		if infile == "" {
			res = uwuify(text)
		} else {
			res = uwuifyFile(infile)
		}

		if outfile == "" {
			fmt.Println(res)
		} else {
			writeToFile(outfile, res)
		}

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

func uwuify(s string) string {
	replaceText(&s, .5)
	addKaomoji(&s, 0.025)
	addExclamations(&s, .5)
	addActions(&s, .025)

	return s
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

func uwuifyFile(file string) string {
	content, err := os.ReadFile(file)
	if err != nil {
		log.Fatalf("error reading file: %s - %s", file, err)
	}
	return string(uwuify(string(content)))
}

func writeToFile(path string, content string) {
	err := os.WriteFile(path, []byte(content), 0644)
	if err != nil {
		log.Fatalf("error writing file: %s - %s", path, err)
	}
}
