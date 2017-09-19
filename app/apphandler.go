package app

type AppHandler struct{
	mux *chi.mux
	//redirectURL string
	session *session
	Logger *log.Logger
}

func NewAppHandler () *AppHandler{
	h := &AppHandler{
		mux: chi.NewRouter(),
		Session: s,
		Logger: log.New(os.Stderr, "", log.LstdFlags),
	}
	h.mux.Get("/app/index", h.indexHandler)
	return h
}
func (h *AppHandler) indexHandler (w http.ResponseWriter, r *http.Request){
	tmpl := template.Must(template.ParseFiles("/templates/analytic.html"))
}