package cmd

import (
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/vikbert/random/utils"
	"regexp"
	"strings"
)

var (
	rootCmd = &cobra.Command{
		Use:   "random",
		Short: "A simple CLI random content generator",
		Long: `A Simple and Flexible Random Content Generator
    built with love with spf13/cobra in Go.
    As default, it generates a random UUID and pastes the UUID to clipboard
    Complete documentation is available at https://github.com/vikbert/random`,
		SilenceErrors: false,
		SilenceUsage:  true,
	}
)

func ExecuteAll() error {
	// hide the unused commmand "completion"
	rootCmd.CompletionOptions.DisableDefaultCmd = true

	// styling via UsageTemplate
	cobra.AddTemplateFunc("StyleHeading", color.New(color.FgGreen).SprintFunc())
	usageTemplate := rootCmd.UsageTemplate()
	usageTemplate = strings.NewReplacer(
		`Usage:`, `{{StyleHeading "Usage:"}}`,
		`Aliases:`, `{{StyleHeading "Aliases:"}}`,
		`Available Commands:`, `{{StyleHeading "Available Commands:"}}`,
		`Global Flags:`, `{{StyleHeading "Global Flags:"}}`,
	).Replace(usageTemplate)
	re := regexp.MustCompile(`(?m)^Flags:\s*$`)
	usageTemplate = re.ReplaceAllLiteralString(usageTemplate, `{{StyleHeading "Flags:"}}`)
	rootCmd.SetUsageTemplate(usageTemplate)

	// apply default command to generate UUID
	id := utils.GenerateUuid()
	utils.ClipCopy(id, false)

	return rootCmd.Execute()
}
