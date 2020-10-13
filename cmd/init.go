package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

func createGgfFile() error {
	file, err := os.Create("ggf.yaml")
	if err != nil {
		return err
	}

	if _, err = fmt.Fprintln(file, "# GottaGoFast"); err != nil {
		return err
	}

	return nil
}

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a project",
	Run: func(cmd *cobra.Command, args []string) {
		if err := createGgfFile(); err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}

		fmt.Println("ggf.yaml created!")
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
