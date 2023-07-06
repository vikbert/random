package commands

import (
	"github.com/spf13/cobra"
	"github.com/vikbert/random/pkg/Generator"
	"github.com/vikbert/random/utils"
)

func init() {
	rootCmd.AddCommand(loremCmd)
	loremCmd.PersistentFlags().IntP(
		"size",
		"s",
		120,
		"Amount of words for generated text",
	)
}

var (
	loremCmd = &cobra.Command{
		Use:   "text",
		Short: "Generate dummy lorem text",
		Long: `This command generates the dummy lorem text 
    and pastes the content to the system clipboard automatically`,

		Run: func(cmd *cobra.Command, args []string) {
			size, _ := cmd.Flags().GetInt("size")

			var loremText string
			loremText = Generator.RandomText(size)
			utils.ClipCopy(loremText, true)
		},
	}
)
