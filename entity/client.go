package entity

import (
	"github.com/docker/docker/client"
)

func GetClient() (*client.Client) {
    client, err := client.NewClientWithOpts()
    if err != nil {
        panic(err)
    }

    return client
}
