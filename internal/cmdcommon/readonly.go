package cmdcommon

import (
	"github.com/spf13/cobra"

	"github.com/ankitpokhrel/jira-cli/internal/cmdutil"
)

func DisableWriteCommand(cmd *cobra.Command) {
	cmd.Hidden = true
	cmd.PersistentPreRun = func(cmd *cobra.Command, args []string) {
		cmdutil.Failed("This command is disabled in readonly mode.")
	}
}
