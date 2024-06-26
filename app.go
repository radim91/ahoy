package main

import (
	"html/template"
	"net/http"
	"os"
	"radim91/entity"
)

func projectsHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFS(files, "templates/project/list.html", "templates/base.html")

	if err != nil {
		panic(err)
	}

	Projects := entity.GetProjects()

	tmpl.ExecuteTemplate(w, "base", Projects)
}

func projectDetailHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFS(files, "templates/project/detail.html", "templates/base.html")

	if err != nil {
		panic(err)
	}

    Data := map[string]interface{} {
        "Project": entity.GetProject(r.PathValue("name")),
        "Url": os.Getenv("URL"),
    }

	tmpl.ExecuteTemplate(w, "base", Data)
}

func containersHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFS(files, "templates/container/list.html", "templates/base.html")

	if err != nil {
		panic(err)
	}

	Containers := entity.GetContainers()

	tmpl.ExecuteTemplate(w, "base", Containers)
}

func imagesHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFS(files, "templates/image/list.html", "templates/base.html")

	if err != nil {
		panic(err)
	}

	Images := entity.GetImages()

	tmpl.ExecuteTemplate(w, "base", Images)
}

func imageDetailHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFS(files, "templates/image/detail.html", "templates/base.html")

	if err != nil {
		panic(err)
	}

	Image := entity.GetImage(r.PathValue("id"))

	tmpl.ExecuteTemplate(w, "base", Image)
}

func networksHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFS(files, "templates/network/list.html", "templates/base.html")

	if err != nil {
		panic(err)
	}

	Networks := entity.GetNetworks()

	tmpl.ExecuteTemplate(w, "base", Networks)
}

func networkDetailHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFS(files, "templates/network/detail.html", "templates/base.html")

	if err != nil {
		panic(err)
	}

	Network := entity.GetNetwork(r.PathValue("id"))
    Containers := entity.GetContainers();

    var data = map[string]interface{} {
        "Network": Network,
        "Containers": Containers,
    }

	tmpl.ExecuteTemplate(w, "base", data)
}

func volumesHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFS(files, "templates/volume/list.html", "templates/base.html")

	if err != nil {
		panic(err)
	}

	Volumes := entity.GetVolumes()

	tmpl.ExecuteTemplate(w, "base", Volumes)
}

func volumeDetailHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFS(files, "templates/volume/detail.html", "templates/base.html")

	if err != nil {
		panic(err)
	}

	volume := entity.GetVolume(r.PathValue("name"))

	tmpl.ExecuteTemplate(w, "base", volume)
}

func containerDetailHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFS(files, "templates/container/detail.html", "templates/base.html")

	if err != nil {
		panic(err)
	}

    Data := map[string]interface{} {
        "Container": entity.GetContainer(r.PathValue("id")),
        "Url": os.Getenv("URL"),
    }

	tmpl.ExecuteTemplate(w, "base", Data)
}

func containerLogsHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFS(files, "templates/container/log.html", "templates/base.html")

	if err != nil {
		panic(err)
	}

    data := map[string]interface{} {
        "Container": entity.GetContainer(r.PathValue("id")),
        "Url": os.Getenv("URL"),
    }

	tmpl.ExecuteTemplate(w, "base", data)
}

func containerStatsHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFS(files, "templates/container/stats.html", "templates/base.html")

	if err != nil {
		panic(err)
	}

    data := map[string]interface{} {
        "Container": entity.GetContainer(r.PathValue("id")),
        "Url": os.Getenv("URL"),
    }

	tmpl.ExecuteTemplate(w, "base", data)
}
