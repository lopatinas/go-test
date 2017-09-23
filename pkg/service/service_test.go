package service

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/lopatinas/go-test/pkg/config"
	"github.com/lopatinas/go-test/pkg/handlers"
	"github.com/lopatinas/go-test/pkg/router/bitroute"
)

func TestSetup(t *testing.T) {
	cfg := new(config.Config)
	err := cfg.Load(config.SERVICENAME)
	if err != nil {
		t.Error("Expected loading of environment vars, got", err)
	}
	router, logger, err := Setup(cfg)
	if err != nil {
		t.Errorf("Fail, got '%s', want '%v'", err, nil)
	}
	if router == nil {
		t.Error("Expected new router, got nil")
	}
	if logger == nil {
		t.Error("Expected new logger, got nil")
	}

	h := handlers.New(logger, cfg)
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		h.Base(notFound)(bitroute.NewControl(w, r))
	})

	req, err := http.NewRequest("GET", "/notfound", nil)
	if err != nil {
		t.Error(err)
	}

	trw := httptest.NewRecorder()
	handler.ServeHTTP(trw, req)

	if trw.Code != http.StatusNotFound {
		t.Error("Expected status:", http.StatusNotFound, "got", trw.Code)
	}
}
