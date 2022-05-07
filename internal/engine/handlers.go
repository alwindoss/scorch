package engine

import (
	"fmt"
	"net/http"
	"time"
)

type ScorchHandler interface {
	Install(w http.ResponseWriter, r *http.Request)
	Status(w http.ResponseWriter, r *http.Request)
	Ping(w http.ResponseWriter, r *http.Request)
}

func NewScorchHandler() ScorchHandler {
	return &scorchHandler{}
}

type scorchHandler struct {
}

func (s *scorchHandler) Install(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Install invoked at %v", time.Now())
}

func (s *scorchHandler) Status(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Status invoked at %v", time.Now())
}

func (s *scorchHandler) Ping(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Ping invoked at %v", time.Now())
}
