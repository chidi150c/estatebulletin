package http

import (
	"html/template"
	"net/http"
	"strings"

	"github.com/pkg/errors"
)

var (
	indexTmpl = newAppTemplate("index.html")
	//IndexTmpl     = newAppTemplate("index.html")
	//LoginTmpl     = newAppTemplate("login.html")
	//ListusersTmpl = newAppTemplate("listusers.html")
	socialTmpl = newAppTemplate("social.html")
)

// appTemplate is a user login-aware wrapper for a html/template.
type appTemplate struct {
	t *template.Template
}

// parseTemplate applies a given file to the body of the base template.
func newAppTemplate(filename string) *appTemplate {
	path := strings.Join([]string{"templates", filename}, "/")
	tmpl := template.Must(template.ParseFiles("templates/base.html", path, "templates/chat.html", "templates/chatlog.html"))
	return &appTemplate{t: tmpl}
}

// func (tmpl *appTemplate) ParseChatTemplate(filename string) (*Template, error) {
// 	func (t *Template) ParseFiles(filenames ...string) (*Template, error)
// }

// Execute writes the template using the provided data, adding login and user
// information to the base template...   usr interface{}, noFooter bool
func (tmpl *appTemplate) Execute(w http.ResponseWriter, r *http.Request, dat interface{}) error {
	d := struct {
		Data        interface{}
		AuthEnabled bool
		LoginURL    string
		LogoutURL   string
		//AddFooter   bool
		SignupURL string
		//User      interface{}
	}{
		Data:        dat,
		AuthEnabled: true,
		LoginURL:    "/login?redirect=" + r.URL.RequestURI(),
		LogoutURL:   "/logout?redirect=" + r.URL.RequestURI(),
		SignupURL:   "/signup?redirect=" + r.URL.RequestURI(),
		//AddFooter:   noFooter,
		//User: usr,
	}

	if err := tmpl.t.Execute(w, d); err != nil {
		return errors.Wrapf(err, "could not write template: %+v", err)
	}
	return nil
}
