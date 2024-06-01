package cmd

import (
	"go-uwu/internal"
	"io"
	"log"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

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

		internal.Uwuify(reader, writer)
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
