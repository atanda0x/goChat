package main

import (
	"html/template"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	file := http.FileServer(http.Dir("/public"))
	mux.Handle("/static/", http.StripPrefix("/static/", file))

	mux.HandleFunc("/", index)
	mux.HandleFunc("/err", err)

	mux.HandleFunc("/login", login)
	mux.HandleFunc("/logout", logout)
	mux.HandleFunc("/signup", signup)
	mux.HandleFunc("/signp_account", signp_account)
	mux.HandleFunc("/auth", auth)

	mux.HandleFunc("/thread/new", newThread)
	mux.HandleFunc("/thread/create", createThread)
	mux.HandleFunc("/thread/post", postThread)
	mux.HandleFunc("/thread/read", readThread)

	server := &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: mux,
	}

	server.ListenAndServe()
}

func index(w http.ResponseWriter, r *http.Request) {
	thread, err := data.Threads()
	if err == nil {
		_, err := session(w, r)

		public_temp_file := []string{
			"templates/layout.html",
			"templates/navbar.html",
			"templates/index.html",
		}

		private_temp_file := []string{
			"templates/layout.html",
			"templates/navbar.html",
			"templates/index.html",
		}

		var templates *template.Template
		if err != nil {
			templates = template.Must(templates.ParseFiles(private_temp_file...))
		} else {
			templates = template.Must(template.ParseFiles(public_temp_file...))
		}

		templates.ExecuteTemplate(w, "layout", threads)
	}
}
