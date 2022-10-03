package server

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/kumojin/k8s-ingress-api/pkg/k8s"
	"github.com/kumojin/k8s-ingress-api/pkg/network"
	echo "github.com/labstack/echo/v4"
	echomiddleware "github.com/labstack/echo/v4/middleware"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
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
	s.EchoServer.Add(http.MethodPost, "/:namespace/ingress", s.createIngress)
	s.EchoServer.Add(http.MethodDelete, "/:namespace/ingress/:name", s.deleteIngress)
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

func (s *server) createIngress(c echo.Context) error {
	namespace := c.Param("namespace")

	dryRun, _ := strconv.ParseBool(c.QueryParam("dryRun"))

	opts := new(k8s.IngressCreateTrimOptions)
	if err := c.Bind(opts); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	config, err := k8s.BuildDefaultKubeConfig()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	k8sClient, err := kubernetes.NewForConfig(config)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	createOpts := metav1.CreateOptions{}
	if dryRun {
		createOpts.DryRun = []string{"All"}
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3000*time.Millisecond)
	defer cancel()

	ingresses := k8sClient.NetworkingV1().Ingresses(namespace)
	ingress, err := ingresses.Create(ctx, k8s.BuildIngressCreateConfig(opts), createOpts)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusCreated, ingress)

	return nil
}

func (s *server) deleteIngress(c echo.Context) error {
	namespace := c.Param("namespace")
	name := c.Param("name")

	dryRun, _ := strconv.ParseBool(c.QueryParam("dryRun"))

	config, err := k8s.BuildDefaultKubeConfig()
	if err != nil {
		return err
	}

	k8sClient, err := kubernetes.NewForConfig(config)
	if err != nil {
		return err
	}

	deleteOpts := metav1.DeleteOptions{}
	if dryRun {
		deleteOpts.DryRun = []string{"All"}
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3000*time.Millisecond)
	defer cancel()

	ingresses := k8sClient.NetworkingV1().Ingresses(namespace)
	err = ingresses.Delete(ctx, name, deleteOpts)
	if err != nil {
		return err
	}

	c.JSON(http.StatusCreated, true)

	return nil
}
