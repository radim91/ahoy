package main

import (
	"fmt"
	"html/template"
	"net/http"
	"radim91/entity"
)

func projectsHandler(w http.ResponseWriter, r *http.Request) {
    tmpl, err := template.New("").ParseFiles("templates/projects.html", "templates/base.html")
    
    if err != nil {
        panic(err)
    }

    Projects := entity.GetProjects()

    tmpl.ExecuteTemplate(w, "base", Projects)
}

func containersHandler(w http.ResponseWriter, r *http.Request) {
    tmpl, err := template.New("").ParseFiles("templates/containers.html", "templates/base.html")
    
    if err != nil {
        panic(err)
    }

    Containers := entity.GetContainers()

    tmpl.ExecuteTemplate(w, "base", Containers)
}

func imagesHandler(w http.ResponseWriter, r *http.Request) {
    tmpl, err := template.New("").ParseFiles("templates/images.html", "templates/base.html")
    
    if err != nil {
        panic(err)
    }

    Images := entity.GetImages()

    tmpl.ExecuteTemplate(w, "base", Images)
}

func networksHandler(w http.ResponseWriter, r *http.Request) {
    tmpl, err := template.New("").ParseFiles("templates/networks.html", "templates/base.html")
    
    if err != nil {
        panic(err)
    }

    Networks := entity.GetNetworks()

    tmpl.ExecuteTemplate(w, "base", Networks)
}

func volumesHandler(w http.ResponseWriter, r *http.Request) {
    tmpl, err := template.New("").ParseFiles("templates/volumes.html", "templates/base.html")
    
    if err != nil {
        panic(err)
    }

    Volumes := entity.GetVolumes()

    tmpl.ExecuteTemplate(w, "base", Volumes)
}

func containerDetailHandler(w http.ResponseWriter, r *http.Request) {
    tmpl, err := template.New("").ParseFiles("templates/container.html", "templates/base.html")
    
    fmt.Println(r.PathValue("id"))
    if err != nil {
        panic(err)
    }

    Container := entity.GetContainer(r.PathValue("id"))

    tmpl.ExecuteTemplate(w, "base", Container)
}

func main() {
    mux := http.NewServeMux()
    mux.HandleFunc("/", projectsHandler)
    mux.HandleFunc("/projects", projectsHandler)
    mux.HandleFunc("/containers", containersHandler)
    mux.HandleFunc("/images", imagesHandler)
    mux.HandleFunc("/networks", networksHandler)
    mux.HandleFunc("/volumes", volumesHandler)

    mux.HandleFunc("/containers/{id}", containerDetailHandler)

    mux.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

    http.ListenAndServe(":8080", mux)
}
