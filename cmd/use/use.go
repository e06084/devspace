package use

import (
	"github.com/devspace-cloud/devspace/cmd/flags"
	"github.com/devspace-cloud/devspace/pkg/devspace/plugin"
	"github.com/devspace-cloud/devspace/pkg/util/factory"
	"github.com/spf13/cobra"
)

// NewUseCmd creates a new cobra command for the use sub command
func NewUseCmd(f factory.Factory, globalFlags *flags.GlobalFlags, plugins []plugin.Metadata) *cobra.Command {
	useCmd := &cobra.Command{
		Use:   "use",
		Short: "Use specific config",
		Long: `
#######################################################
#################### devspace use #####################
#######################################################
	`,
		Args: cobra.NoArgs,
	}

	useCmd.AddCommand(newProfileCmd(f))
	useCmd.AddCommand(newContextCmd(f, globalFlags))
	useCmd.AddCommand(newNamespaceCmd(f, globalFlags))

	// Add plugin commands
	plugin.AddPluginCommands(useCmd, plugins, "use")
	return useCmd
}
