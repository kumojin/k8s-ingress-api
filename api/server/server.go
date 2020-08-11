package server

import (
	"net/http"

	echo "github.com/labstack/echo/v4"
	echomiddleware "github.com/labstack/echo/v4/middleware"
)

func NewServer() (*echo.Echo, error) {

	server := echo.New()

	server.Use(echomiddleware.CORSWithConfig(echomiddleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodOptions, http.MethodGet, http.MethodPost},
	}))

	server.Add(http.MethodGet, "ping", func(ctx echo.Context) error {
		return ctx.String(http.StatusOK, "pong")
	})

	return server, nil
}
