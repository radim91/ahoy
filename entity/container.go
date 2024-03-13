package entity

import (
	"bufio"
	"context"
	"time"

	"github.com/docker/docker/api/types"
)

func GetContainers() []types.Container {
	client := GetClient()

	containers, err := client.ContainerList(context.Background(), types.ContainerListOptions{All: true})
	if err != nil {
		panic(err)
	}

	return containers
}

func GetContainer(id string) types.ContainerJSON {
	client := GetClient()

	container, err := client.ContainerInspect(context.Background(), id)
	if err != nil {
		panic(err)
	}

	return container
}

func GetContainerLogs(id string) []string {
	client := GetClient()
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	logs, _ := client.ContainerLogs(ctx, id, types.ContainerLogsOptions{
		ShowStdout: true,
		ShowStderr: true,
		Details:    true,
		Timestamps: true,
		Since:      "1440m",
	})

	defer logs.Close()

	var logsSlice []string
	scanner := bufio.NewScanner(logs)
	for scanner.Scan() {
		logsSlice = append(logsSlice, scanner.Text()[8:])
	}

	return logsSlice
}

func RemoveContainer(id string) {
	client := GetClient()
	err := client.ContainerRemove(context.Background(), id, types.ContainerRemoveOptions{
		RemoveVolumes: true,
		Force:         true,
	})

	if err != nil {
		panic(err)
	}
}
