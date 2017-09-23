package handlers

import (
	"net/http"
	"testing"

	"github.com/lopatinas/go-test/pkg/config"
	"github.com/lopatinas/go-test/pkg/logger"
	"github.com/lopatinas/go-test/pkg/logger/standard"
	"github.com/lopatinas/go-test/pkg/router/bitroute"
)

func TestReady(t *testing.T) {
	h := New(standard.New(&logger.Config{}), new(config.Config))
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		h.Base(h.Ready)(bitroute.NewControl(w, r))
	})

	testHandler(t, handler, http.StatusOK, http.StatusText(http.StatusOK))
}
