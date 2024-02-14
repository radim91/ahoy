package entity

import (
	"context"

	"github.com/docker/docker/api/types/volume"
)

func GetVolumes() (volume.ListResponse) {
    client := GetClient()

    volumes, err := client.VolumeList(context.Background(), volume.ListOptions{})
    if err != nil {
        panic(err)
    }

    return volumes
}
