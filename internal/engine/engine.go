package engine

import (
	"fmt"
	"log"
	"net/http"

	"github.com/alwindoss/scorch"
	"github.com/gorilla/mux"
)

func Run(cfg *scorch.Config) error {
	router := NewRouter(cfg)
	handler := NewScorchHandler()
	SetupRoutes(cfg, router, handler)
	addr := fmt.Sprintf(":%d", cfg.Port)
	log.Printf("Starting scorch server, listening on port %d", cfg.Port)
	err := http.ListenAndServe(addr, router)
	if err != nil {
		return err
	}
	return nil
}

func NewRouter(cfg *scorch.Config) *mux.Router {
	return mux.NewRouter()
}

func SetupRoutes(cfg *scorch.Config, r *mux.Router, h ScorchHandler) {
	v1Router := r.PathPrefix("/scorch/api/v1").Subrouter()
	v1Router.Path("/ping").Methods(http.MethodGet).HandlerFunc(h.Ping)
	v1Router.Path("/install").Methods(http.MethodPost).HandlerFunc(h.Install)
	v1Router.Path("/status").Methods(http.MethodGet).HandlerFunc(h.Status)
}
