/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
)

// docsCmd represents the docs command
var docsCmd = &cobra.Command{
	Use:   "docs",
	Short: "generates markdown docs",
	RunE: func(cmd *cobra.Command, args []string) error {
		err := doc.GenMarkdownTree(rootCmd, "docs")
		if err != nil {
			return err
		}
		return nil
	},
	Hidden:                true,
	SilenceUsage:          true,
	Args:                  cobra.NoArgs,
	DisableFlagsInUseLine: true,
}

func init() {
	rootCmd.AddCommand(docsCmd)
}
