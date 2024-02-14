package entity

import (
    "context"

    "github.com/docker/docker/api/types"
)

func GetNetworks() ([]types.NetworkResource) {
    client := GetClient()

    networks, err := client.NetworkList(context.Background(), types.NetworkListOptions{})
    if err != nil {
        panic(err)
    }

    return networks
}
