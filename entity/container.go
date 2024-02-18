package entity

import (
	"bufio"
	"context"

	"github.com/docker/docker/api/types"
)

func GetContainers() ([]types.Container) {
    client := GetClient()
    
    containers, err := client.ContainerList(context.Background(), types.ContainerListOptions{All: true})
    if err != nil {
        panic(err)
    }

    return containers
}

func GetContainer(id string) (types.ContainerJSON) {
    client := GetClient()
    
    container, err := client.ContainerInspect(context.Background(), id)
    if err != nil {
        panic(err)
    }

    return container
}

func GetContainerLogs(id string) []string {
    client := GetClient()
    
    logs, _ := client.ContainerLogs(context.Background(), id, types.ContainerLogsOptions{
        ShowStdout: true,
        ShowStderr: true,
        Details: true,
        Timestamps: true,
    })

    defer logs.Close()

    var logsSlice []string
    scanner := bufio.NewScanner(logs)
    for scanner.Scan() {
        logsSlice = append(logsSlice, scanner.Text())
    }

    return logsSlice
}
