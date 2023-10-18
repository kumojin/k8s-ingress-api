package handler

import (
	"encoding/json"
	"github.com/kumojin/k8s-ingress-api/api/config"
	"github.com/kumojin/k8s-ingress-api/pkg/k8s"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"os"
	"strconv"
)

type handler struct {
	client    *k8s.Client
	config    config.IngressConfig
	errLogger *log.Logger
	logger    *log.Logger
}

func NewHandler(kc *k8s.Client, ic config.IngressConfig) handler {
	return handler{kc, ic, log.New(os.Stderr, "", 0), log.New(os.Stdout, "", 0)}
}

func (h *handler) CreateIngress(c echo.Context) error {
	dryRun, _ := strconv.ParseBool(c.QueryParam("dryRun"))
	host := c.QueryParam("host")

	ingress, err := h.client.CreateIngress(host, dryRun)
	if err != nil {
		h.errLogger.Printf("Could not create ingress: %v\n", err.Error())
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	i, _ := json.Marshal(ingress)
	h.logger.Printf("Created ingress: %v\n", string(i))
	return c.JSON(http.StatusCreated, ingress)
}

func (h *handler) DeleteIngress(c echo.Context) error {
	//name := c.Param("name")
	//
	//dryRun, _ := strconv.ParseBool(c.QueryParam("dryRun"))
	//host := c.QueryParam("host")
	//
	//deleteOpts := metav1.DeleteOptions{}
	//if dryRun {
	//	deleteOpts.DryRun = []string{"All"}
	//}
	//
	//ctx, cancel := context.WithTimeout(context.Background(), 3000*time.Millisecond)
	//defer cancel()
	//
	//ingresses := h.client.Client.NetworkingV1().Ingresses(h.config.Namespace)
	//err := ingresses.Delete(ctx, name, deleteOpts)
	//if err != nil {
	//	return err
	//}
	//
	//c.JSON(http.StatusCreated, true)

	return nil
}
