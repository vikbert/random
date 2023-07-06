package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/vikbert/random/utils"
)

func init() {
	rootCmd.AddCommand(jokeCmd)
}

var (
	jokeCmd = &cobra.Command{
		Use:   "joke",
		Short: "Fetch a random joke from joke-API",
		Long: `This command fetches a random joke from a free joke-api 
    and displays the content in the terminal`,

		Run: func(_ *cobra.Command, args []string) {
			fmt.Println("\n", utils.FetchJoke())
		},
	}
)
