/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "ptm",
	Short: "cmd helper",
	Long:  `A cmd helper for translate and manage papers.`,

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("	WELCOME TO ptm,A CMD HELPER FOR TRANSLATE AND MANAGE PAPERS ON LINUX!")
		fmt.Println("USAGE:")
		fmt.Println("	ptm t     translate papers      ")
		fmt.Println("	ptm m     manage papers         ")

	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(t)

}
