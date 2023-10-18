package server

import (
	"github.com/kumojin/k8s-ingress-api/api/config"
	"github.com/kumojin/k8s-ingress-api/api/handler"
	"github.com/kumojin/k8s-ingress-api/pkg/k8s"
	"github.com/kumojin/k8s-ingress-api/pkg/network"
	echo "github.com/labstack/echo/v4"
	echomiddleware "github.com/labstack/echo/v4/middleware"
	"net/http"
)

type server struct {
	EchoServer       *echo.Echo
	ingressConfig    config.IngressConfig
	kubernetesConfig config.KubernetesConfig
}

type validateCNAMEResponse struct {
	CNAME   string `json:"cname"`
	Matches string `json:"matches"`
	Ok      bool   `json:"ok"`
}

func NewServer(kubernetesConfig config.KubernetesConfig, ingressConfig config.IngressConfig, kclient *k8s.Client) *server {
	s := &server{
		ingressConfig:    ingressConfig,
		kubernetesConfig: kubernetesConfig,
	}

	s.EchoServer = echo.New()
	s.attachHandlers(kclient)
	return s
}

func (s *server) attachHandlers(kc *k8s.Client) {
	s.EchoServer.Use(echomiddleware.CORSWithConfig(echomiddleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodOptions, http.MethodGet, http.MethodPost},
	}))

	s.EchoServer.Add(http.MethodGet, "/ping", s.ping)
	s.EchoServer.Add(http.MethodGet, "/cname/:cname/matches/:matches", s.validateCNAME)

	h := handler.NewHandler(kc, s.ingressConfig)
	s.EchoServer.Add(http.MethodPost, "/ingress", h.CreateIngress)
	//s.EchoServer.Add(http.MethodDelete, "/ingress/:name", h.DeleteIngress)
}

func (s *server) Start(port string) error {
	return s.EchoServer.Start(port)
}

func (s *server) ping(c echo.Context) error {
	return c.JSON(http.StatusOK, "pong")
}

func (s *server) validateCNAME(c echo.Context) error {
	cname := c.Param("cname")
	matches := c.Param("matches")

	ok, err := network.ValidateCNAME(cname, matches)
	if err != nil {
		return err
	}

	c.JSON(http.StatusOK, &validateCNAMEResponse{
		CNAME:   cname,
		Matches: matches,
		Ok:      ok,
	})

	return nil
}
