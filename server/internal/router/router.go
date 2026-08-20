package router

import (
	"net/http"

	"github.com/iamneek/linkpie/internal/handler"
)

func New() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /health", handler.HealthChk)
	return mux
}
