package init

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

func createGgfFile() error {
	file, err := os.Create("ggf.json")
	if err != nil {
		return err
	}

	if _, err = fmt.Fprintln(file, "# GottaGoFast"); err != nil {
		return err
	}

	return nil
}

func NewInitCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "init",
		Short: "Initialize a project",
		Run: func(cmd *cobra.Command, args []string) {
			if err := createGgfFile(); err != nil {
				fmt.Println("Error:", err)
				os.Exit(1)
			}

			fmt.Println("ggf.json created!")
		},
	}
}
