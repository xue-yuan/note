package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add note",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if tag == "" {
			tag = "general"
		}

		stringArgs := strings.Join(args, " ")
		if v, ok := noteMap[tag]; ok {
			noteMap[tag] = append(v.([]interface{}), stringArgs)
		} else {
			noteMap[tag] = []string{stringArgs}
		}

		jsonData, err := json.Marshal(noteMap)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		jsonFile, err := os.OpenFile(filePath, os.O_WRONLY, 0755)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		defer jsonFile.Close()

		_, err = jsonFile.Write(jsonData)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
