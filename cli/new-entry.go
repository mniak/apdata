package main

import (
	"github.com/mniak/apdata"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(&newEntryCmd)
}

var newEntryCmd = cobra.Command{
	Use: "new-entry",
	Aliases: []string{
		"add-entry",
		"mark-entry",
		"entry",
	},
	Short: "Creates a new timecard entry",
	Run: func(cmd *cobra.Command, args []string) {
		company, err := cmd.Flags().GetString("company")
		handle(err)
		username, err := cmd.Flags().GetString("user")
		handle(err)
		password, err := cmd.Flags().GetString("password")
		handle(err)

		verbose, err := cmd.Flags().GetBool("verbose")
		handle(err)

		baseurl := getBaseUrl(company)
		client, err := apdata.NewClient(baseurl,
			apdata.WithDebug(verbose),
		)
		handle(err)

		err = client.Login(username, password)
		handle(err)

		err = client.MarkEntry()
		handle(err)
	},
}
