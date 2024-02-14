package entity

import (
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

