package cmd

import (
	"go-uwu/internal"
	"io"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var (
	text    string
	infile  string
	outfile string

	kaomoji          float64
	textReplacements float64
	stutters         float64
	exclamations     float64
	actions          float64
	allowUnicode     bool

	rootCmd = &cobra.Command{
		Use:   "uwuify",
		Short: "uwuifies the given text",
		Long: `transforms the given input and outputs it to the desired destination
e.g. uwuify --infile input.txt -o ~/output.txt -u 1

input can be defined via --infile, --text, or stdin via the "|" operator
the transformed text will be output to stdout by default or written to a file by using --outfile

the modifiers --uwu, --kaomoji, --stutters, --exclamations, and --actions can be used to set the probability for the corresponding transforms
0  -> will not occur
.5 -> will occur 50% of the time
1  -> will occur at every possibility, usually for, or after every word`,
		SilenceUsage: true,
		Args:         cobra.NoArgs,
		RunE:         runFunc,
		Version:      "0.1",
	}
)

func runFunc(cmd *cobra.Command, args []string) error {
	var reader io.Reader

	stat, _ := os.Stdin.Stat()

	switch {
	case stat.Mode()&os.ModeDevice == 0:
		// if the input does not come from a character device (e.g. a terminal), it is most likely piped in
		reader = cmd.InOrStdin()
	case infile != "":
		file, err := os.Open(infile)
		if err != nil {
			return err
		}
		defer file.Close()
		reader = file
	case text != "":
		reader = strings.NewReader(text)
	default:
		cmd.Help()
		return nil
	}

	var writer io.Writer

	if outfile != "" {
		file, err := os.Create(outfile)
		if err != nil {
			return err
		}
		defer file.Close()
		writer = file
	} else {
		writer = cmd.OutOrStdout()
	}

	return internal.Uwuify(reader, writer, internal.Options{
		TextReplacements: textReplacements,
		Stutters:         stutters,
		Kaomoji:          kaomoji,
		Exclamations:     exclamations,
		Actions:          actions,
		Unicode:          allowUnicode,
	})
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.CompletionOptions.HiddenDefaultCmd = true
	rootCmd.PersistentFlags().SortFlags = false
	rootCmd.Flags().SortFlags = false

	rootCmd.PersistentFlags().StringVarP(&text, "text", "t", "", "a text to uwuify")
	rootCmd.PersistentFlags().StringVarP(&infile, "infile", "i", "", "a file to uwuify")
	rootCmd.PersistentFlags().StringVarP(&outfile, "outfile", "o", "", "a file to output the uwuified text to")
	rootCmd.MarkFlagsMutuallyExclusive("text", "infile")

	rootCmd.PersistentFlags().BoolVarP(&allowUnicode, "unicode", "u", false, `allow unicode characters for kaomoji`)

	rootCmd.PersistentFlags().Float64VarP(&textReplacements, "replacements", "r", .5, `probability for transforming text. e.g. love -> wuv`)
	rootCmd.PersistentFlags().Float64VarP(&kaomoji, "kaomoji", "k", .025, `probability for inserting kaomoji. e.g. OwO`)
	rootCmd.PersistentFlags().Float64VarP(&stutters, "stutters", "s", .025, `probability for adding stutters to the beginning of a word. e.g. hello -> h-hello`)
	rootCmd.PersistentFlags().Float64VarP(&exclamations, "exclamations", "e", .5, `probability for transforming punctuation. e.g. 1 -> !!11`)
	rootCmd.PersistentFlags().Float64VarP(&actions, "actions", "a", .025, `probability for adding actions. e.g. *blushes*`)
	rootCmd.Aliases = append(rootCmd.Aliases, "uwu")

}
