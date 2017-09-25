package http

import (
	"net/http"
	"strings"
)

type HttpHandler struct {
	appHandler *AppHandler
}

func NewHttpHandler(a *AppHandler) *HttpHandler {
	return &HttpHandler{
		appHandler: a,
	}
}

func (h *HttpHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	if strings.HasPrefix(r.URL.Path, "/asset/") {
		http.StripPrefix("/asset/", http.FileServer(http.Dir("./asset/"))).ServeHTTP(w, r)
	} else if strings.HasPrefix(r.URL.Path, "/") {
		h.appHandler.ServeHTTP(w, r)
	}
}
