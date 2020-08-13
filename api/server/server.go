package server

import (
	"net/http"

	"github.com/kumojin/k8s-ingress-api/pkg/network"
	echo "github.com/labstack/echo/v4"
	echomiddleware "github.com/labstack/echo/v4/middleware"
)

type server struct {
	EchoServer *echo.Echo
}

type validateCNAMEResponse struct {
	CNAME   string `json:"cname"`
	Matches string `json:"matches"`
	Ok      bool   `json:"ok"`
}

func NewServer() *server {
	s := &server{}
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
	//TODO create ingress
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
