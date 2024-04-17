package main

import (
	"net/http"

	"github.com/joho/godotenv"
)

func init() {
    err := godotenv.Load(".env")
    if err != nil {
        panic(err)
    }
}

func main() {
	// app.go
	mux := http.NewServeMux()
	mux.HandleFunc("/", projectsHandler)
	mux.HandleFunc("/projects", projectsHandler)
	mux.HandleFunc("/containers", containersHandler)
	mux.HandleFunc("/images", imagesHandler)
	mux.HandleFunc("/networks", networksHandler)
	mux.HandleFunc("/volumes", volumesHandler)

	mux.HandleFunc("/containers/{id}", containerDetailHandler)
	mux.HandleFunc("/containers/{id}/logs", containerLogsHandler)
	mux.HandleFunc("/images/{id}", imageDetailHandler)
	mux.HandleFunc("/projects/{name}", projectDetailHandler)
	mux.HandleFunc("/networks/{id}", networkDetailHandler)
	mux.HandleFunc("/volumes/{name}", volumeDetailHandler)

	// api.go
	mux.HandleFunc("/api/container/start/{id}", apiContainerStartHandler)
	mux.HandleFunc("/api/container/stop/{id}", apiContainerStopHandler)
	mux.HandleFunc("/api/container/logs/{id}", apiContainerLogsHandler)
	mux.HandleFunc("/api/container/status/{id}", apiContainerStatusHandler)
	mux.HandleFunc("/api/container/remove/{id}", apiContainerRemoveHandler)
	mux.HandleFunc("/api/image/remove/{id}", apiImageRemoveHandler)
	mux.HandleFunc("/api/project/start/{name}", apiProjectStartHandler)
	mux.HandleFunc("/api/project/stop/{name}", apiProjectStopHandler)
	mux.HandleFunc("/api/project/restart/{name}", apiProjectRestartHandler)
	mux.HandleFunc("/api/project/down/{name}", apiProjectDownHandler)
	mux.HandleFunc("/api/project/status/{name}", apiProjectStatusHandler)
    mux.HandleFunc("/api/volume/remove/{name}", apiVolumeRemoveHandler)
	mux.HandleFunc("POST /api/network/connect", apiNetworkAddContainerHandler)
	mux.HandleFunc("POST /api/network/disconnect", apiNetworkRemoveContainerHandler)

	// FS
	mux.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

	http.ListenAndServe(":8080", mux)
}
