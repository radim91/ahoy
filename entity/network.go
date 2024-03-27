package entity

import (
	"context"
	"fmt"

	"github.com/docker/docker/api/types"
)

func GetNetworks() []types.NetworkResource {
	client := GetClient()

	networks, err := client.NetworkList(context.Background(), types.NetworkListOptions{})
	if err != nil {
		panic(err)
	}

	return networks
}

func GetNetwork(id string) types.NetworkResource {
	client := GetClient()

	network, err := client.NetworkInspect(context.Background(), id, types.NetworkInspectOptions{})
	if err != nil {
		panic(err)
	}

	return network
}

func AddContainerToNetwork(id string, containerId string) {
    client := GetClient()

    err := client.NetworkConnect(context.Background(), id, containerId, nil)

    if err != nil {
        fmt.Println(err)
    }
}

func RemoveContainerFromNetwork(id string, containerId string) {
    client := GetClient()

    err := client.NetworkDisconnect(context.Background(), id, containerId, true)

    if err != nil {
        fmt.Println(err)
    }
}
