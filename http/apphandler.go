package http

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
)

type AppHandler struct {
	mux *chi.Mux
	//redirectURL string
	//session *session
	Logger *log.Logger
}

//pass this as parameter when session is implemented: (s *Session)
func NewAppHandler() *AppHandler {
	h := &AppHandler{
		mux: chi.NewRouter(),
		//Session: s,
		Logger: log.New(os.Stderr, "", log.LstdFlags),
	}
	h.mux.Get("/app/0", h.indexHandler)
	h.mux.Get("/app/1", h.socialHandler)
	// h.mux.Get("/app/3", h.apartmentHandler)
	// h.mux.Get("/app/4", h.hotelHandler)
	// h.mux.Get("/app/5", h.analyticHandler)
	return h
}

func (a *AppHandler) indexHandler(w http.ResponseWriter, r *http.Request) {
	if err := indexTmpl.t.Execute(w, nil); err != nil {
		log.Fatal(err)
	}
}
func (a *AppHandler) socialHandler(w http.ResponseWriter, r *http.Request) {
	if err := socialTmpl.t.Execute(w, nil); err != nil {
		log.Fatal(err)
	}
}

// func (a *AppHandler) apartmentHandler(w http.ResponseWriter, r *http.Request) {
// 	tmpl := template.Must(template.ParseFiles("templates/apartment.html"))
// 	if err := tmpl.Execute(w, nil); err != nil {
// 		log.Fatal(err)
// 	}
// }
// func (a *AppHandler) hotelHandler(w http.ResponseWriter, r *http.Request) {
// 	tmpl := template.Must(template.ParseFiles("templates/hotel.html"))
// 	if err := tmpl.Execute(w, nil); err != nil {
// 		log.Fatal(err)
// 	}
// }
// func (a *AppHandler) analyticHandler(w http.ResponseWriter, r *http.Request) {
// 	tmpl := template.Must(template.ParseFiles("templates/analytic.html"))
// 	if err := tmpl.Execute(w, nil); err != nil {
// 		log.Fatal(err)
// 	}
// }
func (a *AppHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.mux.ServeHTTP(w, r)
}
