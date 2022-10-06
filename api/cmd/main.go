package cmd

import (
  "github.com/kumojin/k8s-ingress-api/api/server"
  "github.com/spf13/viper"

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
  viper.AutomaticEnv()

  viper.SetConfigType("yaml")
  viper.SetConfigName("config")
  viper.AddConfigPath("/etc/k8s-ingress-api/config")
  viper.AddConfigPath(".")

  viper.ReadInConfig()

  s := server.NewServer()
  s.EchoServer.Logger.Fatal(s.Start(":3000"))
}
