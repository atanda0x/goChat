package main

import (
	"fmt"
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

func index(writer http.ResponseWriter, request *http.Request) {
	threads, err := data.Threads()
	if err == nil {
		_, err := session(writer, request)
		if err != nil {
			generateHTML(writer, threads, "layout", "public.navbar", "index")
		} else {
			generateHTML(writer, threads, "layout", "private.navbar", "index")
		}
	}
}

func generateHTML(w http.ResponseWriter, data interface{}, fn ...string) {
	var files []string

	for _, file := range fn {
		files = append(files, fmt.Sprintf("templates/%s.html", file))

	}
	template := template.Must(template.ParseFiles(files...))
	template.ExecuteTemplate(w, "layout", data)
}
