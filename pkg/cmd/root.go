package cmd

import (
	generateCommand "github.com/ThebaultLouis/GottaGoFast/pkg/cmd/generate"
	initCommand "github.com/ThebaultLouis/GottaGoFast/pkg/cmd/init"
	"github.com/spf13/cobra"
)

func NewRootCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "ggf",
		Short: "A code generator that doesn't waste time.",
	}

	cmd.AddCommand(initCommand.NewInitCommand())
	cmd.AddCommand(generateCommand.NewGenerateCommand())

	return cmd
}
