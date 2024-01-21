package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path"

	"github.com/spf13/cobra"
)

var tag string
var version bool
var noteMap map[string]interface{}
var filePath string

var rootCmd = &cobra.Command{
	Use:   "note",
	Short: "Note is a quick note tool",
	Long: `Note is a fast and flexible command-line tool to
help you take a note quickly is the terminal.
This application is built with spf13/cobra.`,
	Run: func(cmd *cobra.Command, args []string) {
		if version {
			fmt.Println("Note 1.0")
		}
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	directory := "/.note"
	filename := "note.json"
	filePath = path.Join(home, directory, filename)

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		file, err := os.Create(filePath)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		defer file.Close()

		emptyJSON := map[string]interface{}{}
		emptyJSONBytes, err := json.Marshal(emptyJSON)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		_, err = file.Write(emptyJSONBytes)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		noteMap = emptyJSON
	} else if err != nil {
		fmt.Println(err)
		os.Exit(1)
	} else {
		jsonFile, err := os.Open(filePath)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		defer jsonFile.Close()

		byteValue, err := io.ReadAll(jsonFile)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		json.Unmarshal([]byte(byteValue), &noteMap)
	}

	rootCmd.PersistentFlags().StringVarP(&tag, "tag", "t", "", "tag for note")
	rootCmd.Flags().BoolVarP(&version, "version", "v", false, "print note version")
}
