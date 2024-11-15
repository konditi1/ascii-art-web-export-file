package handler

import (
	"html/template"
	"net/http"

	webart "webart/server"
)

// Asciiweb handles http requests and responses
func Asciiweb(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		if r.Method == http.MethodGet {
			tmpl, err := template.ParseFiles("template/index.html")
			if err != nil {
				webart.ShowError(w, "server error", "Internal Server Error", http.StatusInternalServerError)
				return
			}
			w.WriteHeader(http.StatusOK)
			tmpl.Execute(w, nil)
		} else {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}
	} else if r.URL.Path == "/ascii-art" {
		webart.ProcessForm(w, r, "template/index.html")
	} else if r.URL.Path == "/about" {
		tmpl, err := template.ParseFiles("template/about.html")
		if err != nil {
			webart.ShowError(w, "server error", "Internal Server Error", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		tmpl.Execute(w, nil)
	} else if r.URL.Path == "/save" {
		webart.Save(w,r)
	} else {
		webart.ShowError(w, "page not found", "Page Not Found", http.StatusNotFound)
		return
	}
}
