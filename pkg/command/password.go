package commands

import (
	"github.com/spf13/cobra"
	"github.com/vikbert/random/pkg/Generator"
	"github.com/vikbert/random/utils"
)

func init() {
	rootCmd.AddCommand(passCmd)
}

var (
	passLength = 12
	passCmd    = &cobra.Command{
		Use:   "pass",
		Short: "Generate a random password",
		Long: `This command generates a random password and 
    pastes the password to the system clipboard automatically. Default length of password is 12`,
		Run: func(_ *cobra.Command, _ []string) {
			password := Generator.GeneratePassword(passLength)
			utils.ClipCopy(password, true)

		},
	}
)
