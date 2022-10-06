package server

import (
	"github.com/kumojin/k8s-ingress-api/api/config"
	"github.com/kumojin/k8s-ingress-api/api/handler"
	"github.com/kumojin/k8s-ingress-api/pkg/network"
	echo "github.com/labstack/echo/v4"
	echomiddleware "github.com/labstack/echo/v4/middleware"
	"net/http"
)

type server struct {
	EchoServer    *echo.Echo
	IngressConfig config.IngressConfig
}

type validateCNAMEResponse struct {
	CNAME   string `json:"cname"`
	Matches string `json:"matches"`
	Ok      bool   `json:"ok"`
}

func NewServer() *server {
	s := &server{
		IngressConfig: config.GetIngressConfig(),
	}

	s.EchoServer = echo.New()
	s.attachHandlers()
	return s
}

func (s *server) attachHandlers() {
	s.EchoServer.Use(echomiddleware.CORSWithConfig(echomiddleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodOptions, http.MethodGet, http.MethodPost},
	}))

	s.EchoServer.Add(http.MethodGet, "/ping", s.ping)
	s.EchoServer.Add(http.MethodGet, "/cname/:cname/matches/:matches", s.validateCNAME)

	h := handler.NewHandler(s.IngressConfig)
	s.EchoServer.Add(http.MethodPost, "/:namespace/ingress", h.CreateIngress)
	s.EchoServer.Add(http.MethodDelete, "/:namespace/ingress/:name", h.DeleteIngress)
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
