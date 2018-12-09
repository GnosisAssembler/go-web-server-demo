package main

// “net/http” to access the core go http functionality
// “fmt” for formatting text
// “html/template” a library for interaction with the html file.

import (
	"fmt"
	"html/template"
	"net/http"
)

// Create a struct that holds information to be displayed in the HTML file
type Hello struct {
	Name string
}

// Go app entrypoint
func main() {
	// Init a Hello struct object and pass in some random information.
	hello := Hello{"Anonymous"}

	// Get html template
	templates := template.Must(template.ParseFiles("template/index.html"))

	http.Handle("/static/", //final url can be anything
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("static"))))

	// This method takes in the URL path "/" and a function that takes in a response writer, and a http request.
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		// Takes the name from the URL query e.g ?name=Samy, will set welcome.Name = Samy.
		if name := r.FormValue("name"); name != "" {
			hello.Name = name
		}
		// If errors show an internal server error message
		// I also pass the hello struct to the index.html file.
		if err := templates.ExecuteTemplate(w, "index.html", hello); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	// Start the web server, set the port to listen to 8000
	// Print any errors from starting the webserver using fmt
	fmt.Println("Go Server Listening on port 8000")
	fmt.Println(http.ListenAndServe(":8000", nil))
}
