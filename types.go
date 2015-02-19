package gocluster

type DiscoveryBackend interface {
	Connect(string, *Cluster)
	Register()
}
