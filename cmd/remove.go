package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

var removeCmd = &cobra.Command{
	Use:   "remove <index>",
	Short: "Remove note by specific index",
	Args:  cobra.ExactArgs(1),
	PreRun: func(cmd *cobra.Command, args []string) {
		if tag == "" {
			tag = "general"
		}

		i, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		if _, ok := noteMap[tag]; !ok {
			fmt.Println("The tag is not exist.")
			os.Exit(1)
		}

		if i >= len(noteMap[tag].([]interface{})) {
			fmt.Println("Invalid index.")
			os.Exit(1)
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		if tag == "" {
			tag = "general"
		}

		i, _ := strconv.Atoi(args[0])
		noteMap[tag] = append(noteMap[tag].([]interface{})[:i], noteMap[tag].([]interface{})[i+1:]...)

		jsonData, err := json.Marshal(noteMap)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		jsonFile, err := os.OpenFile(filePath, os.O_WRONLY|os.O_TRUNC, 0755)
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
	rootCmd.AddCommand(removeCmd)
	// removeCmd.
}
