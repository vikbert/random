package cmd

import (
	"github.com/spf13/cobra"
	"github.com/vikbert/random/utils"
)

var (
	baseCmd = &cobra.Command{
		Use:   "random",
		Short: "A simple CLI random content generator",
		Long: `A Simple and Flexible Random Content Generator
    built with love with spf13/cobra in Go.
    Complete documentation is available at https://github.com/vikbert/random`,
		SilenceErrors: true,
		SilenceUsage:  true,
	}
)

func ExecuteAll() error {
	baseCmd.CompletionOptions.DisableDefaultCmd = true

	id := utils.GenerateUuid()
	utils.ClipCopy(id, false)

	return baseCmd.Execute()
}
