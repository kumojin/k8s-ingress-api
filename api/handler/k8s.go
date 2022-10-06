package handler

import (
	"context"
	"github.com/kumojin/k8s-ingress-api/api/config"
	"github.com/kumojin/k8s-ingress-api/pkg/k8s"
	"github.com/labstack/echo/v4"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"net/http"
	"strconv"
	"time"
)

type handler struct {
	client *k8s.Client
	config config.IngressConfig
}

func NewHandler(config config.IngressConfig) handler {
	client, err := k8s.NewClient(config)
	if err != nil {
		panic(err)
	}

	return handler{client, config}
}

func (h *handler) CreateIngress(c echo.Context) error {
	dryRun, _ := strconv.ParseBool(c.QueryParam("dryRun"))

	opts := new(k8s.IngressCreateTrimOptions)
	if err := c.Bind(opts); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	createOpts := metav1.CreateOptions{}
	if dryRun {
		createOpts.DryRun = []string{"All"}
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3000*time.Millisecond)
	defer cancel()

	ingresses := h.client.Client.NetworkingV1().Ingresses(h.config.Namespace)
	ingress, err := ingresses.Create(ctx, k8s.BuildIngressCreateConfig(opts), createOpts)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusCreated, ingress)

	return nil
}

func (h *handler) DeleteIngress(c echo.Context) error {
	name := c.Param("name")

	dryRun, _ := strconv.ParseBool(c.QueryParam("dryRun"))

	deleteOpts := metav1.DeleteOptions{}
	if dryRun {
		deleteOpts.DryRun = []string{"All"}
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3000*time.Millisecond)
	defer cancel()

	ingresses := h.client.Client.NetworkingV1().Ingresses(h.config.Namespace)
	err := ingresses.Delete(ctx, name, deleteOpts)
	if err != nil {
		return err
	}

	c.JSON(http.StatusCreated, true)

	return nil
}
