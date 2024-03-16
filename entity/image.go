package entity

import (
	"context"

	"github.com/docker/docker/api/types"
)

func GetImages() []types.ImageSummary {
	client := GetClient()

	images, err := client.ImageList(context.Background(), types.ImageListOptions{})
	if err != nil {
		panic(err)
	}

	return images
}

func GetImage(id string) types.ImageInspect {
	client := GetClient()

	image, _, err := client.ImageInspectWithRaw(context.Background(), id)
	if err != nil {
		panic(err)
	}

	return image
}

func RemoveImage(id string) []types.ImageDeleteResponseItem {
	client := GetClient()
	response, err := client.ImageRemove(context.Background(), id, types.ImageRemoveOptions{
		Force:         true,
		PruneChildren: true,
	})

	if err != nil {
		panic(err)
	}

	return response
}
