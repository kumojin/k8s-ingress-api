package cmd

import (
	"kumojin/k8s-ingress-api/api/server"

	"github.com/spf13/cobra"
)

func MainCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "web",
		Short: "Starts the Web server",
		Run:   bootWebServer,
	}
}

func bootWebServer(cmd *cobra.Command, args []string) {
	server, err := server.NewServer()
	if err != nil {
		panic(err)
	}
	server.Logger.Fatal(server.Start(":3000"))
}
