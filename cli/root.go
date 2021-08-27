package main

import "github.com/spf13/cobra"

var rootCmd = cobra.Command{
	Use: "apdata",
}

func main() {
	rootCmd.PersistentFlags().String("company", "", "Company identifier on the URL")
	rootCmd.MarkPersistentFlagRequired("company")

	rootCmd.PersistentFlags().StringP("user", "u", "", "Username")
	rootCmd.MarkPersistentFlagRequired("user")

	rootCmd.PersistentFlags().StringP("password", "p", "", "Password")
	rootCmd.MarkPersistentFlagRequired("password")

	rootCmd.PersistentFlags().BoolP("verbose", "v", false, "Prints more detailed information")

	rootCmd.CompletionOptions.DisableDefaultCmd = true
	rootCmd.Execute()
}
