package cmd

import (
	"os"

	"mm-patch/patch"

	"github.com/spf13/cobra"
)

var fname string

var patchCmd = &cobra.Command{
	Use:   "patch",
	Short: "Patch Mattermost binary for license generation",
	RunE: func(cmd *cobra.Command, args []string) error {
		err := patch.ApplyPatch(fname)
		if err != nil {
			cmd.PrintErrln(err)
			os.Exit(1)
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(patchCmd)
	patchCmd.Flags().StringVarP(&fname, "filename", "f", "/opt/mattermost/bin/mattermost", "Path to mattermost binary file")
}
