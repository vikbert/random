package command

import (
	"github.com/spf13/cobra"
	"github.com/vikbert/random/pkg/clipboard"
	"github.com/vikbert/random/pkg/generator"
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
			loremText = generator.RandomText(size)
			clipboard.ClipCopy(loremText, true)
		},
	}
)
