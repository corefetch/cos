package main

import (
	"context"
	"fmt"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/client"
)

func main() {

	apiClient, err := client.NewClientWithOpts(client.FromEnv)

	if err != nil {
		panic(err)
	}
	defer apiClient.Close()

	containers, err := apiClient.ContainerList(context.Background(), container.ListOptions{
		All: true,
		Filters: filters.NewArgs(
			filters.KeyValuePair{
				Key:   "label",
				Value: "com.docker.compose.project",
			},
		),
	})
	if err != nil {
		panic(err)
	}

	for _, ctr := range containers {
		fmt.Printf("%s - %s\n", ctr.ID, ctr.Labels["com.docker.compose.project"])
	}
}
