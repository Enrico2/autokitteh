package connections

import (
	"fmt"

	"github.com/spf13/cobra"

	"go.autokitteh.dev/autokitteh/cmd/ak/common"
	"go.autokitteh.dev/autokitteh/internal/resolver"
	"go.autokitteh.dev/autokitteh/sdk/sdkservices"
)

var listCmd = common.StandardCommand(&cobra.Command{
	Use:     "list [--integration=...] [--fail]",
	Short:   "List all connections",
	Aliases: []string{"ls", "l"},
	Args:    cobra.NoArgs,

	RunE: func(cmd *cobra.Command, args []string) error {
		var f sdkservices.ListConnectionsFilter

		if integration != "" {
			r := resolver.Resolver{Client: common.Client()}
			_, iid, err := r.IntegrationNameOrID(integration)
			if err != nil {
				return err
			}

			if !iid.IsValid() {
				return fmt.Errorf("integration %q not found", integration)
			}
			f.IntegrationID = iid
		}

		ctx, cancel := common.LimitedContext()
		defer cancel()

		cs, err := connections().List(ctx, f)
		if err != nil {
			return fmt.Errorf("list connections: %w", err)
		}

		if err := common.FailIfNotFound(cmd, "connections", len(cs) > 0); err != nil {
			return err
		}

		common.RenderList(cs)
		return nil
	},
})

func init() {
	// Command-specific flags.
	listCmd.Flags().StringVarP(&integration, "integration", "i", "", "integration name or ID")

	common.AddFailIfNotFoundFlag(listCmd)
}
