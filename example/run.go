package main

import (
	"os"

	"../"
)

func main() {
	mongoUrl := getConfig("MONGO_DISCOVERY_URL", "mongodb://localhost/discovery")
	service := getConfig("CLUSTER_SERVICE", "gobackend")
	endpointUrl := getConfig("CLUSTER_ENDPOINT_URL", "http://myip:port")

	cluster := gocluster.Connect(mongoUrl)
	cluster.PingInterval = 1000 * 1 // default 5 sec
	cluster.Register(service, endpointUrl)

	<-make(chan bool)
}

func getConfig(key string, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		value = defaultValue
	}

	return value
}
