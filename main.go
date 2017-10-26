package main

import (
	// in-built libraries
	"path"
	"html/template"
	"log"
	"net/http"

	// third party libraries
	"github.com/gorilla/mux"
)

type Color struct {
	Color_code string
}

// home page handler
func indexHandler(res http.ResponseWriter, req *http.Request) {
	res.WriteHeader(http.StatusOK)
	res.Write([]byte("whatisthiscolor.rexlabs.xyz visually gives output of how would a color look like"))
}

func colorHandler(res http.ResponseWriter, req *http.Request) {
	parameter := mux.Vars(req)
	c := Color{string(parameter["colorcode"])}
	res.WriteHeader(http.StatusOK)
	// log.Printf("Color page Input Color: %v\n", parameter["colorcode"])
	file := path.Join("templates", "index.html")
	tmpl, err := template.ParseFiles(file)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := tmpl.Execute(res, c); err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	// create a router
	r := mux.NewRouter()
	r.HandleFunc("/", indexHandler)
	r.HandleFunc("/{colorcode}", colorHandler)

	// attach the router and serve it
	log.Println("whatisthiscolor.rexlabs.xyz is live on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
