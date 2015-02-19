package main

import (
	"../"
)

func main() {
	cluster := gocluster.Connect("mongodb://localhost/discovery")
	cluster.Register("web", "http://google.com")
}
