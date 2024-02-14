package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"radim91/entity"
)

func projectsHandler(w http.ResponseWriter, r *http.Request) {
    tmpl, err := template.ParseFiles("templates/projects.html")
    
    if err != nil {
        panic(err)
    }

    Projects := entity.GetProjects()
    fmt.Println(Projects)

    tmpl.Execute(w, Projects)
}

func containersHandler(w http.ResponseWriter, r *http.Request) {
    jsonData, err := json.Marshal(entity.GetContainers())
    
    if err != nil {
        panic(err)
    }

    w.Write(jsonData)
}

func main() {
    mux := http.NewServeMux()
    mux.HandleFunc("/", projectsHandler)
    mux.HandleFunc("/projects", projectsHandler)
    mux.HandleFunc("/containers", containersHandler)

    mux.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

    http.ListenAndServe(":8080", mux)
}
