package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"radim91/entity"
    "time"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type NetworkConnectRequest struct {
    ContainerId, NetworkId string
}

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
        time.Sleep(1 * time.Second)
		status := entity.GetContainerStatus(r.PathValue("id"))
        connErr := conn.WriteMessage(websocket.TextMessage, []byte(status))

        if connErr != nil {
            fmt.Println(connErr)

            break
        }
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
        time.Sleep(1 * time.Second)
		logs := entity.GetContainerLogs(r.PathValue("id"))
		jsonData, _ := json.Marshal(logs)
        err := conn.WriteMessage(websocket.TextMessage, jsonData)

        if err != nil {
            fmt.Println(err)

            break
        }
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

func apiImageRemoveHandler(w http.ResponseWriter, r *http.Request) {
	removeConfirmation := entity.RemoveImage(r.PathValue("id"))

	jsonData, _ := json.Marshal(removeConfirmation)

	w.Write(jsonData)
}

func apiNetworkAddContainerHandler(w http.ResponseWriter, r *http.Request) {
    dec := json.NewDecoder(r.Body)
    var body NetworkConnectRequest

    err := dec.Decode(&body)
    if err != nil {
        panic(err)
    }

    entity.AddContainerToNetwork(body.NetworkId, body.ContainerId)

    msg := map[string]string{
        "message": "connected",
    }

    jsonData, _ := json.Marshal(msg)

    w.Write(jsonData)
}

func apiNetworkRemoveContainerHandler(w http.ResponseWriter, r *http.Request) {
    dec := json.NewDecoder(r.Body)
    var body NetworkConnectRequest

    err := dec.Decode(&body)
    if err != nil {
        panic(err)
    }

    entity.RemoveContainerFromNetwork(body.NetworkId, body.ContainerId)

    msg := map[string]string{
        "message": "disconnected",
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
        time.Sleep(1 * time.Second)
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

func apiVolumeRemoveHandler(w http.ResponseWriter, r *http.Request) {
    go entity.RemoveVolume(r.PathValue("name"))

    msg := map[string]string{
        "message": "removed",
    }

    jsonData, _ := json.Marshal(msg)

    w.Write(jsonData)
}
