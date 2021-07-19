package main

import (
	log "github.com/sirupsen/logrus"
	"ns2docker"
)

func main(){
	ns2docker.LoadDockerNsCache()
	log.Infof("查询ns: %v","4026532229")

	dockerContainer := ns2docker.Search("4026532229")
	log.Infof("使用namespace: %v 查询到容器: %v","4026532229",dockerContainer)

	if namespace,err := ns2docker.QueryNs("bf5d97f08b5a");err == nil {
		log.Infof("使用容器ID：%s 查询到namespace: %v","bf5d97f08b5a",namespace)
	}else{
		log.Error(err.Error())
	}

}
