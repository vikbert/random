package cmd

import (
	"github.com/spf13/cobra"
	"github.com/vikbert/random/utils"
	"strconv"
)

func init() {
	baseCmd.AddCommand(loremCmd)
}

var (
	loremCmd = &cobra.Command{
		Use:   "text",
		Short: "Generate dummy lorem text",
		Long: `This command generates the dummy lorem text 
    and pastes the content to the system clipboard automatically`,

		Run: func(_ *cobra.Command, args []string) {
			var count int
			if len(args) >= 1 {
				arg1, err := strconv.Atoi(args[0])
				if err != nil {
					count = 0
				} else {
					count = arg1
				}
			}

			loremText := utils.GenerateText(count)
			utils.ClipCopy(loremText, true)
		},
	}
)
