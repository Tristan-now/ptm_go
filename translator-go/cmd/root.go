package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "ptm_go",
	Short: "Papers translation and management",
	Long:  `A command line application to translate using deepl API and manage papers.`,

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to ptm_go!")
		fmt.Println("-t to translate")

	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
