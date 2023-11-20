package main

import (
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	router := mux.NewRouter()

	// Serve static files from the "static" directory
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Home route
	router.HandleFunc("/", HomeHandler).Methods("GET")

	// Overview route
	router.HandleFunc("/overview", OverviewHandler).Methods("GET")

	// Toggle content route
	router.HandleFunc("/toggle-content", ToggleContentHandler).Methods("GET")

	http.Handle("/", router)
	http.ListenAndServe(":8080", nil)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	name := struct {
		Title       string
		Description string
	}{
		Title:       "Varmint Cong",
		Description: "Simple Go Web Starter Kit",
	}

	// Render the HTML template
	templates.ExecuteTemplate(w, "index.html", name)
}

func OverviewHandler(w http.ResponseWriter, r *http.Request) {
	name := struct {
		Title string
	}{
		Title: "Project Overview",
	}

	// Render the HTML template for the project overview
	templates.ExecuteTemplate(w, "overview.html", name)
}

func ToggleContentHandler(w http.ResponseWriter, r *http.Request) {
	// You can add logic here to determine the content to be toggled
	name := struct {
		Title string
	}{
		Title: "Toggled Content",
	}

	// Render the HTML template for the toggled content
	templates.ExecuteTemplate(w, "toggle-content.html", name)
}
