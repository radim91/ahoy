package entity

import (
    "context"

    "github.com/docker/docker/api/types"
)

func GetImages() ([]types.ImageSummary) {
    client := GetClient()

    images, err := client.ImageList(context.Background(), types.ImageListOptions{})
    if err != nil {
        panic(err)
    }

    return images
}
