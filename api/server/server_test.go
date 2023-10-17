package server

import (
	"github.com/kumojin/k8s-ingress-api/api/config"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPing(t *testing.T) {
	kc := config.GetKubernetesConfig()
	kc.InCluster = false
	ic := config.IngressConfig{}
	s := NewServer(kc, ic)

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()

	c := s.EchoServer.NewContext(req, rec)
	c.SetPath("/ping")

	if assert.NoError(t, s.ping(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "\"pong\"\n", string(rec.Body.Bytes()))
	}
}

func TestValidateCNAMENotOk(t *testing.T) {
	kc := config.GetKubernetesConfig()
	kc.InCluster = false
	ic := config.IngressConfig{}
	s := NewServer(kc, ic)

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()

	c := s.EchoServer.NewContext(req, rec)
	c.SetPath("/cname/:cname/matches/:matches")
	c.SetParamNames("cname", "matches")
	c.SetParamValues("m.facebook.com", "star-mini.facebook.com")

	if assert.NoError(t, s.validateCNAME(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "{\"cname\":\"m.facebook.com\",\"matches\":\"star-mini.facebook.com\",\"ok\":false}\n", string(rec.Body.Bytes()))
	}
}

func TestValidateCNAMEOk(t *testing.T) {
	kc := config.GetKubernetesConfig()
	kc.InCluster = false
	ic := config.IngressConfig{}
	s := NewServer(kc, ic)

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()

	c := s.EchoServer.NewContext(req, rec)
	c.SetPath("/cname/:cname/matches/:matches")
	c.SetParamNames("cname", "matches")
	c.SetParamValues("m.facebook.com", "star-mini.c10r.facebook.com")

	if assert.NoError(t, s.validateCNAME(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "{\"cname\":\"m.facebook.com\",\"matches\":\"star-mini.c10r.facebook.com\",\"ok\":true}\n", string(rec.Body.Bytes()))
	}
}

//TODO create ingress

//TODO delete ingress
