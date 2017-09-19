package app

import "text/template"

type AppHandler struct{
	mux *chi.mux
	//redirectURL string
	session *session
	Logger *log.Logger
}

//pass this as parameter when session is implemented: (s *Session)
func NewAppHandler () *AppHandler{
	h := &AppHandler{
		mux: chi.NewRouter(),
		//Session: s,
		Logger: log.New(os.Stderr, "", log.LstdFlags),
	}
	h.mux.Get("/app/index", h.indexHandler)
	return h
}
func (a *AppHandler) indexHandler (w http.ResponseWriter, r *http.Request){
	tmpl := template.Must(template.ParseFiles("/templates/analytic.html"))
	if err := tmpl.Execute(w, nil); err!=nil{
		log.Fatal(err)
	}
}

func (a *AppHandler) serveHTTP (w http.ResponseWriter, r *http.Request){
	a.mux.serveHTTP(w,r)
}
	