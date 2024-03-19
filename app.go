package main

import (
	"html/template"
	"net/http"
	"radim91/entity"
)

func projectsHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.New("").ParseFiles("templates/project/list.html", "templates/base.html")

	if err != nil {
		panic(err)
	}

	Projects := entity.GetProjects()

	tmpl.ExecuteTemplate(w, "base", Projects)
}

func projectDetailHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.New("").ParseFiles("templates/project/detail.html", "templates/base.html")

	if err != nil {
		panic(err)
	}

	Project := entity.GetProject(r.PathValue("name"))

	tmpl.ExecuteTemplate(w, "base", Project)
}

func containersHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.New("").ParseFiles("templates/container/list.html", "templates/base.html")

	if err != nil {
		panic(err)
	}

	Containers := entity.GetContainers()

	tmpl.ExecuteTemplate(w, "base", Containers)
}

func imagesHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.New("").ParseFiles("templates/image/list.html", "templates/base.html")

	if err != nil {
		panic(err)
	}

	Images := entity.GetImages()

	tmpl.ExecuteTemplate(w, "base", Images)
}

func imageDetailHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.New("").ParseFiles("templates/image/detail.html", "templates/base.html")

	if err != nil {
		panic(err)
	}

	Image := entity.GetImage(r.PathValue("id"))

	tmpl.ExecuteTemplate(w, "base", Image)
}

func networksHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.New("").ParseFiles("templates/network/list.html", "templates/base.html")

	if err != nil {
		panic(err)
	}

	Networks := entity.GetNetworks()

	tmpl.ExecuteTemplate(w, "base", Networks)
}

func networkDetailHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.New("").ParseFiles("templates/network/detail.html", "templates/base.html")

	if err != nil {
		panic(err)
	}

	Network := entity.GetNetwork(r.PathValue("id"))

	tmpl.ExecuteTemplate(w, "base", Network)
}

func volumesHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.New("").ParseFiles("templates/volume/list.html", "templates/base.html")

	if err != nil {
		panic(err)
	}

	Volumes := entity.GetVolumes()

	tmpl.ExecuteTemplate(w, "base", Volumes)
}

func volumeDetailHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.New("").ParseFiles("templates/volume/detail.html", "templates/base.html")

	if err != nil {
		panic(err)
	}

	volume := entity.GetVolume(r.PathValue("name"))

	tmpl.ExecuteTemplate(w, "base", volume)
}

func containerDetailHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.New("").ParseFiles("templates/container/detail.html", "templates/base.html")

	if err != nil {
		panic(err)
	}

	Container := entity.GetContainer(r.PathValue("id"))

	tmpl.ExecuteTemplate(w, "base", Container)
}

func containerLogsHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.New("").ParseFiles("templates/container/log.html", "templates/base.html")

	if err != nil {
		panic(err)
	}

	tmpl.ExecuteTemplate(w, "base", r.PathValue("id"))
}
