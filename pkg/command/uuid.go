package command

import (
	"github.com/spf13/cobra"
	"github.com/vikbert/random/pkg/clipboard"
	"github.com/vikbert/random/pkg/generator"
)

func init() {
	rootCmd.AddCommand(uuidCmd)
}

var (
	uuidCmd = &cobra.Command{
		Use:   "uuid",
		Short: "Generate a random UUID",
		Long: `This command generates a random UUID 
    and pastes the UUID to the system clipboard automatically`,

		Run: func(_ *cobra.Command, _ []string) {
			id := generator.GenerateUuid()
			clipboard.ClipCopy(id, true)
		},
	}
)
