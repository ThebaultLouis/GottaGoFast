package cmd

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"path"
)

func findGgfFileDirectory() (string, error) {
	workingDirectory, err := os.Getwd()
	if err != nil {
		return "", err
	}

	// We look for the file ggf.yaml starting from the working directory
	currentDirectory := workingDirectory

	for currentDirectory != "/" {
		ggfFilePath := path.Join(currentDirectory, "ggf.yaml")

		if _, err := os.Stat(ggfFilePath); err == nil {
			// We found the file!
			return currentDirectory, nil
		}

		// We look at the parent directory
		currentDirectory = path.Dir(currentDirectory)
	}

	// The file does not exist!
	return "", errors.New("ggf.yaml not found")
}

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Do the thing",
	Run: func(cmd *cobra.Command, args []string) {
		ggfFileDirectory, err := findGgfFileDirectory()
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}

		fmt.Println("Found ggf.yaml at", ggfFileDirectory)
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)
}
