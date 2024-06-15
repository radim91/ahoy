package main

import (
	"embed"
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

var (
	//go:embed templates/*
	files embed.FS
	//go:embed assets/*
	assets embed.FS
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		os.Setenv("URL", "localhost:8080")
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
	mux.HandleFunc("/containers/{id}/stats", containerStatsHandler)
	mux.HandleFunc("/images/{id}", imageDetailHandler)
	mux.HandleFunc("/projects/{name}", projectDetailHandler)
	mux.HandleFunc("/networks/{id}", networkDetailHandler)
	mux.HandleFunc("/volumes/{name}", volumeDetailHandler)

	// api.go
	mux.HandleFunc("/api/container/start/{id}", apiContainerStartHandler)
	mux.HandleFunc("/api/container/stop/{id}", apiContainerStopHandler)
	mux.HandleFunc("/api/container/logs/{id}", apiContainerLogsHandler)
	mux.HandleFunc("/api/container/stats/{id}", apiContainerStatsHandler)
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
	mux.Handle("/assets/", http.FileServer(http.FS(assets)))

	fmt.Println("Listening on " + os.Getenv("URL"))
	http.ListenAndServe(os.Getenv("URL"), mux)
}
