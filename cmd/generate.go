package cmd

import (
	"fmt"
	"os"

	"mm-patch/patch"

	"github.com/spf13/cobra"
)

var licCustom patch.LicenseCustomize

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate license file for patched Mattermost",
	RunE: func(cmd *cobra.Command, args []string) error {

		licFilename := fmt.Sprintf("%s.mattermost-license", licCustom.Company)
		err := patch.GenerateLicense(licCustom, licFilename)
		if err != nil {
			cmd.PrintErrln(err)
			os.Exit(1)
		}
		cmd.Printf("License file successfully generated: %s\n", licFilename)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)

	generateCmd.Flags().IntVarP(&licCustom.ExpireYear, "expire-year", "y", 2050, "License expire year")
	generateCmd.Flags().IntVarP(&licCustom.Users, "users", "u", 10000, "License users count")
	generateCmd.Flags().StringVarP(&licCustom.Name, "name", "n", "acme", "License customer name")
	generateCmd.Flags().StringVarP(&licCustom.Email, "email", "e", "test@test.com", "License customer email")
	generateCmd.Flags().StringVarP(&licCustom.Company, "company", "c", "ACME", "License customer company")
}
