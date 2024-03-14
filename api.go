package main

import (
	"encoding/json"
	"log"
	"net/http"
	"radim91/entity"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

func apiContainerStartHandler(w http.ResponseWriter, r *http.Request) {
	go entity.StartContainer(r.PathValue("id"))

	msg := map[string]string{
		"message": "started",
	}

	jsonData, _ := json.Marshal(msg)

	w.Write(jsonData)
}

func apiContainerStatusHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade failed: ", err)
		return
	}

	defer conn.Close()

	for {
		status := entity.GetContainerStatus(r.PathValue("id"))
		conn.WriteMessage(websocket.TextMessage, []byte(status))
	}
}

func apiContainerStopHandler(w http.ResponseWriter, r *http.Request) {
	go entity.StopContainer(r.PathValue("id"))

	msg := map[string]string{
		"message": "stopped",
	}

	jsonData, _ := json.Marshal(msg)

	w.Write(jsonData)
}

func apiContainerLogsHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade failed: ", err)
		return
	}

	defer conn.Close()

	for {
		logs := entity.GetContainerLogs(r.PathValue("id"))
		jsonData, _ := json.Marshal(logs)
		conn.WriteMessage(websocket.TextMessage, jsonData)
	}
}

func apiContainerRemoveHandler(w http.ResponseWriter, r *http.Request) {
	entity.RemoveContainer(r.PathValue("id"))

	msg := map[string]string{
		"message": "removed",
	}

	jsonData, _ := json.Marshal(msg)

	w.Write(jsonData)
}

func apiProjectStartHandler(w http.ResponseWriter, r *http.Request) {
	go entity.StartProject(entity.GetProject(r.PathValue("name")))

	msg := map[string]string{
		"message": "started",
	}

	jsonData, _ := json.Marshal(msg)

	w.Write(jsonData)
}

func apiProjectStopHandler(w http.ResponseWriter, r *http.Request) {
	go entity.StopProject(entity.GetProject(r.PathValue("name")))

	msg := map[string]string{
		"message": "stopped",
	}

	jsonData, _ := json.Marshal(msg)

	w.Write(jsonData)
}

func apiProjectRestartHandler(w http.ResponseWriter, r *http.Request) {
	go entity.RestartProject(entity.GetProject(r.PathValue("name")))

	msg := map[string]string{
		"message": "restarting",
	}

	jsonData, _ := json.Marshal(msg)

	w.Write(jsonData)
}

func apiProjectStatusHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade failed: ", err)
		return
	}

	defer conn.Close()

	for {
		status := entity.GetProjectStatus(r.PathValue("name"))
		conn.WriteMessage(websocket.TextMessage, []byte(status))
	}
}

func apiProjectDownHandler(w http.ResponseWriter, r *http.Request) {
	go entity.DownProject(entity.GetProject(r.PathValue("name")))

	msg := map[string]string{
		"message": "down",
	}

	jsonData, _ := json.Marshal(msg)

	w.Write(jsonData)
}
