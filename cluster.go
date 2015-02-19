package gocluster

import (
	"fmt"
)

type Cluster struct {
	ServiceName  string
	EndpointUrl  string
	EndpointHash string
	PingInterval int64

	discovery DiscoveryBackend
}

func (c *Cluster) Register(serviceName string, endpointUrl string) {
	c.ServiceName = serviceName
	c.EndpointUrl = endpointUrl
	c.PingInterval = 1000 * 1
	c.EndpointHash = ToSha1(endpointUrl)

	fmt.Println("Cluster: Service registered as: ", serviceName)
	fmt.Println("Cluster: Endpoint URL is: ", endpointUrl)

	c.discovery.Register()
}

func Connect(mongoUrl string) *Cluster {
	cluster := &Cluster{}
	cluster.discovery = &MongoDiscovery{}
	cluster.discovery.Connect(mongoUrl, cluster)

	return cluster
}
