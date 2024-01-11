package cmd

import (
	"github.com/spf13/cobra"
	"gohub/pkg/console"
	"gohub/pkg/helpers"
)

var CmdKey = &cobra.Command{
	Use:   "key",
	Short: "Generate App Key, will print the generated key",
	Run:   runKeyGenerate,
	Args:  cobra.NoArgs,
}

func runKeyGenerate(cmd *cobra.Command, args []string) {
	console.Success("---")
	console.Success("App Key:")
	console.Success(helpers.RandomString(32))
	console.Success("---")
	console.Success("please go to .env file to change the APP_KEY option")
}
