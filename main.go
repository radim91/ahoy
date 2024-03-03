package main

import (
	"net/http"
)

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
	mux.HandleFunc("/projects/{name}", projectDetailHandler)

	// api.go
	mux.HandleFunc("/api/container/logs/{id}", apiContainerLogsHandler)
	mux.HandleFunc("/api/project/start/{name}", apiProjectStartHandler)
	mux.HandleFunc("/api/project/stop/{name}", apiProjectStopHandler)
	mux.HandleFunc("/api/project/restart/{name}", apiProjectRestartHandler)
	mux.HandleFunc("/api/project/status/{name}", apiProjectStatusHandler)

	// FS
	mux.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

	http.ListenAndServe(":8080", mux)
}
