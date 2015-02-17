package gocluster

import (
	"fmt"
)

type Cluster struct {
	mongoUrl string
}

func (c *Cluster) Register(serviceName string) {
	fmt.Println("Cluster: Service registered as: ", serviceName)
}

func Connect(mongoUrl string) *Cluster {
	cluster := &Cluster{mongoUrl}
	return cluster
}
