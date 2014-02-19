package main

import (
    "fmt"
    "log"
    "net/http"
    "html/template"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Method, r.URL, "-", r.UserAgent(), "-", r.RemoteAddr)
    fmt.Fprintf(w, "Hello there!")
}

func pageHandler(w http.ResponseWriter, r *http.Request) {
    t, err := template.ParseFiles("./template.html")

    if err != nil {
    	log.Fatal(err)
    }

    toTemplate := struct {
        Name string
        Surname string
    }{
        "Steve",
        "Jobs",
    }
    
    t.Execute(w, toTemplate)
}

func main() {
    http.HandleFunc("/", rootHandler)
    http.HandleFunc("/page", pageHandler)
	
	log.Println("Starting server...")
    log.Fatal(http.ListenAndServe(":80", nil))
}