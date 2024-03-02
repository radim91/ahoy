package main

import (
	"encoding/json"
	"log"
	"net/http"
	"radim91/entity"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

func apiContainerLogsHandler(w http.ResponseWriter, r *http.Request) {
	logs := entity.GetContainerLogs(r.PathValue("id"))
	jsonData, _ := json.Marshal(logs)

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

func apiProjectLogsHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade failed: ", err)
		return
	}

	defer conn.Close()

	for {
		logs := entity.GetProjectLogs(r.PathValue("name"))
		jsonData, _ := json.Marshal(logs)
		conn.WriteMessage(websocket.TextMessage, jsonData)
	}
	/* logs := entity.GetProjectLogs(r.PathValue("name")) */
	/* jsonData, _ := json.Marshal(logs) */
	/**/
	/* w.Write(jsonData) */
}

func apiProjectStatusHandler(w http.ResponseWriter, r *http.Request) {
	status := entity.GetProjectStatus(r.PathValue("name"))
	jsonData, _ := json.Marshal(status)

	w.Write(jsonData)
}
