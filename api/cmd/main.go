package cmd

import (
	"github.com/kumojin/k8s-ingress-api/api/config"
	"github.com/kumojin/k8s-ingress-api/api/server"
	"github.com/kumojin/k8s-ingress-api/pkg/k8s"
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
	viper.AddConfigPath("/etc/k8s-ingress-api")
	viper.AddConfigPath(".")

	viper.ReadInConfig()

	kc := config.GetKubernetesConfig()
	ic := config.GetIngressConfig()
	kclient, err := k8s.NewClient(kc, ic)
	if err != nil {
		panic(err)
	}

	s := server.NewServer(kc, ic, kclient)
	s.EchoServer.Logger.Fatal(s.Start(":3000"))
}
