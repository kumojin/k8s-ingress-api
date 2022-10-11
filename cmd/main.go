package cmd

import (
	apiCmd "github.com/kumojin/k8s-ingress-api/api/cmd"

	"github.com/spf13/cobra"
)

var (
	cfgFile string
)

func RootCommand() *cobra.Command {
	rootCommand := &cobra.Command{Use: "App"}

	rootCommand.AddCommand(apiCmd.MainCommand())

	return rootCommand
}
