package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var showCmd = &cobra.Command{
	Use:   "show",
	Short: "Show note",
	Run: func(cmd *cobra.Command, args []string) {
		tag, err := cmd.Flags().GetString("tag")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		if len(noteMap) == 0 {
			fmt.Println("There is no content in your note.")
			os.Exit(0)
		}

		if tag == "" {
			for k, v := range noteMap {
				fmt.Printf("[%s]\n", k)
				for i, v := range v.([]interface{}) {
					fmt.Printf("%d: %s\n", i, v)
				}
			}
		} else {
			for i, v := range noteMap[tag].([]interface{}) {
				fmt.Printf("%d: %s\n", i, v)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(showCmd)
}
