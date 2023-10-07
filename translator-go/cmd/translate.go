package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"translator-go/deepl"
	"translator-go/utils"
)

func init() {
	rootCmd.AddCommand(t)
}

var t = &cobra.Command{
	Use:   "translate",
	Short: "Translate the selected text",
	Run: func(cmd *cobra.Command, args []string) {
		DEEPL_TOKEN := "ff5294c4-2874-22da-c873-5419f08c2b08:fx"
		for {
			output, err := utils.Get_select_text()
			if output == "" {
				continue
			}
			re_output := utils.Regular_output(output)
			if err != nil {
				fmt.Println(err)
			}
			re, err := deepl.Translate(re_output, "zh", "en", false, DEEPL_TOKEN)
			if err != nil {
				fmt.Println("Please install xsel")
			}
			fmt.Println(output)
			fmt.Println()
			fmt.Println(re)
			fmt.Println()
			fmt.Println("---------------------------------------------")
			fmt.Println()
			/*			if i == 2 {
						cmd := exec.Command("clear") // Unix/Linux/macOS 清屏命令
						cmd.Stdout = os.Stdout
						cmd.Run()
						i = 0
					}*/
		}
	},
}
