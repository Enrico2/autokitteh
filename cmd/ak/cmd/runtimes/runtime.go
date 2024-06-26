package runtimes

import (
	"github.com/spf13/cobra"

	backendRuntimes "go.autokitteh.dev/autokitteh/backend/runtimes"
	"go.autokitteh.dev/autokitteh/cmd/ak/common"
	"go.autokitteh.dev/autokitteh/sdk/sdkservices"
)

// Flag shared by all subcommands.
var local bool

var runtimeCmd = common.StandardCommand(&cobra.Command{
	Use:     "runtime",
	Short:   "Runtime engines: build, get, list, run",
	Aliases: []string{"rt"},
	Args:    cobra.NoArgs,
})

// AddSubcommands adds this command, and its own subcommands, to the calling parent.
func AddSubcommands(parentCmd *cobra.Command) {
	parentCmd.AddCommand(runtimeCmd)
}

func init() {
	// Flag shared by all subcommands.
	runtimeCmd.PersistentFlags().BoolVarP(&local, "local", "l", false, "execute locally")

	// Subcommands.
	runtimeCmd.AddCommand(buildCmd)
	runtimeCmd.AddCommand(getCmd)
	runtimeCmd.AddCommand(listCmd)
	runtimeCmd.AddCommand(runCmd)
	runtimeCmd.AddCommand(testCmd)
}

func runtimes() sdkservices.Runtimes {
	if local {
		return backendRuntimes.New()
	}
	return common.Client().Runtimes()
}
