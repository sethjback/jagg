package cmd

import "github.com/spf13/cobra"

func New() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "jagg",
		Short: "Just Another Guessing Game",
		Run:   func(cmd *cobra.Command, args []string) { cmd.Help() },
	}

	cmd.AddCommand(NewGame())

	return cmd
}
