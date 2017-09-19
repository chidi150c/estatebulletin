package http_api

type Handler struct {
	appHandler app.AppHandler
}

func NewMyHandler(a *app.AppHandler) *MyHandler{
	return &MyHandler{
		appHandler : a
		
	}
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request){
	if strings.HasPrefix(r.URL.Path, "/asset/"){
		http.StripPrefix("/asset/", http.FileServer(http.Dir("./asset/"))).ServeHTTP(w, r)
	}else if strings.HasPrefix(r.URL.Path, "/app"){
		h.appHandler.ServeHTTP(w,r)
	}
}