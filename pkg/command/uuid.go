package commands

import (
	"github.com/spf13/cobra"
	"github.com/vikbert/random/pkg/Generator"
	"github.com/vikbert/random/utils"
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
			id := Generator.GenerateUuid()
			utils.ClipCopy(id, true)
		},
	}
)
