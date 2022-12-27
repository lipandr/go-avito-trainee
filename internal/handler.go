package internal

import (
	"html/template"
	"log"
	"net/http"
)

type formData struct {
	*UserData
	Errors []string
}

var baseLayout = "layout.html"
var templateNames = []string{
	"form.html",
	"thanks.html",
	"sorry.html",
	"list.html",
}
var templates = make(map[string]*template.Template, 3)

func LoadHTMLTemplates() {
	for id, name := range templateNames {
		t, err := template.ParseFiles(baseLayout, name)
		if err != nil {
			log.Fatal(err)
		}
		templates[name] = t
		log.Printf("Template id:%d loaded %s", id, name)
	}
}

func (srv *server) ListHandler(w http.ResponseWriter, _ *http.Request) {
	data := srv.svc.getAllData()
	_ = templates["list.html"].Execute(w, data)
}

func (srv *server) FormHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		_ = templates["form.html"].Execute(w, formData{
			UserData: &UserData{}, Errors: []string{},
		})
	case http.MethodPost:
		err := r.ParseForm()
		if err != nil {
			return
		}
		formFields := UserData{
			Email:      r.Form["email"][0],
			URLtoTrack: r.Form["url"][0],
		}
		var errors []string
		if formFields.Email == "" {
			errors = append(errors, "Please enter your email address")
		}
		if formFields.URLtoTrack == "" {
			errors = append(errors, "Please enter Avito URL to track price")
		}
		if len(errors) > 0 {
			_ = templates["form.html"].Execute(w, formData{
				UserData: &formFields, Errors: errors,
			})
			return
		}
		srv.svc.saveData(formFields)

		if formFields.Email == "1" {
			_ = templates["thanks.html"].Execute(w, formFields.URLtoTrack)
		} else {
			_ = templates["sorry.html"].Execute(w, formFields.URLtoTrack)
		}

	}
}
