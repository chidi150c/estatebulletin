package http

import (
	"net/http"
	"strings"

	"github.com/estatebulletin/app"
)

type HttpHandler struct {
	appHandler *app.AppHandler
}

func NewHttpHandler(a *app.AppHandler) *HttpHandler {
	return &HttpHandler{
		appHandler: a,
	}
}

func (h *HttpHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if strings.HasPrefix(r.URL.Path, "/asset/") {
		http.StripPrefix("/asset/", http.FileServer(http.Dir("./asset/"))).ServeHTTP(w, r)
	} else if strings.HasPrefix(r.URL.Path, "/app") {
		h.appHandler.ServeHTTP(w, r)
	}
}
