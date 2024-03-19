package entity

import (
	"context"

	"github.com/docker/docker/api/types/volume"
)

type Volume struct {
	ClusterVolume *volume.ClusterVolume
	CreatedAt     string
	Driver        string
	Labels        map[string]string
	Mountpoint    string
	Name          string
	Options       map[string]string
	Scope         string
	Status        map[string]interface{}
	UsageData     *volume.UsageData
}

func GetVolumes() map[string]Volume {
	client := GetClient()

	volumes, err := client.VolumeList(context.Background(), volume.ListOptions{})
	if err != nil {
		panic(err)
	}

	volumesMap := make(map[string]Volume)

	for _, volume := range volumes.Volumes {
		volumesMap[volume.Name] = Volume{
			ClusterVolume: volume.ClusterVolume,
			CreatedAt:     volume.CreatedAt,
			Driver:        volume.Driver,
			Labels:        volume.Labels,
			Mountpoint:    volume.Mountpoint,
			Name:          volume.Name,
			Options:       volume.Options,
			Scope:         volume.Scope,
			Status:        volume.Status,
			UsageData:     volume.UsageData,
		}
	}

	return volumesMap
}

func GetVolume(name string) volume.Volume {
	client := GetClient()

	volume, err := client.VolumeInspect(context.Background(), name)
	if err != nil {
		panic(err)
	}

	return volume
}
