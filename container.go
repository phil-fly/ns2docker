package ns2docker

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"os"
)

func LoadDockerNsCache() {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return
	}

	DockerNsCache.Clear()

	containers, err := cli.ContainerList(ctx, types.ContainerListOptions{})
	if err != nil {
		return
	}

	var containerInspect types.ContainerJSON
	for _, container := range containers {
		if containerInspect,err = cli.ContainerInspect(ctx, container.ID);err != nil {
			continue
		}

		filelink, err := os.Readlink(fmt.Sprintf("/proc/%d/ns/pid", containerInspect.State.Pid))
		if err != nil {
			continue
		}

		DockerNsCache.Put(filelink[5:len(filelink)-1],containerInspect)
	}
}

func AddDockerNsCache() {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return
	}

	containers, err := cli.ContainerList(ctx, types.ContainerListOptions{})
	if err != nil {
		return
	}

	var containerInspect types.ContainerJSON
	for _, container := range containers {
		if containerInspect,err = cli.ContainerInspect(ctx, container.ID);err != nil {
			continue
		}

		filelink, err := os.Readlink(fmt.Sprintf("/proc/%d/ns/pid", containerInspect.State.Pid))
		if err != nil {
			continue
		}

		DockerNsCache.Put(filelink[5:len(filelink)-1],containerInspect)
	}
}

func QueryNs(containerID string) (namespace string,err error) {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return "",err
	}

	var containerInspect types.ContainerJSON
	if containerInspect,err = cli.ContainerInspect(ctx, containerID);err != nil {
		return "",err
	}

	filelink, err := os.Readlink(fmt.Sprintf("/proc/%d/ns/pid", containerInspect.State.Pid))
	if err != nil {
		return "",err
	}

	return filelink[5:len(filelink)-1],nil
}

func SearchContainerName(namespace string) string {

	if namespace == "4026531836" {
		return "localhost"
	}

	//直接获取失败时，新增加载ns列表
	dockerContainer,ok := DockerNsCache.Get(namespace)
	if ok {
		return dockerContainer.Name[1:]
	}

	AddDockerNsCache()
	dockerContainer,ok = DockerNsCache.Get(namespace)
	if ok {
		return dockerContainer.Name[1:]
	}
	return ""
}

func SearchOverlay2(namespace string,typeName string) string {

	if namespace == "4026531836" {
		return ""
	}

	//直接获取失败时，新增加载ns列表
	dockerContainer,ok := DockerNsCache.Get(namespace)
	if ok {
		return  dockerContainer.GraphDriver.Data[typeName]
	}

	AddDockerNsCache()
	dockerContainer,ok = DockerNsCache.Get(namespace)
	if ok {
		return dockerContainer.GraphDriver.Data[typeName]
	}
	return ""
}