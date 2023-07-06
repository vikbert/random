package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/vikbert/random/utils"
)

func init() {
	rootCmd.AddCommand(mottoCmd)
}

var (
	mottoCmd = &cobra.Command{
		Use:   "motto",
		Short: "Fetch a random motto from API Ninjas",
		Long: `This command fetches a random motto from a free quotes-api powered by API Ninjas 
    and displays the content in the terminal`,

		Run: func(_ *cobra.Command, args []string) {
			fmt.Println("\n", utils.FetchMotto())
		},
	}
)
