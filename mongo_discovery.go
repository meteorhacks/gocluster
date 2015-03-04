package gocluster

import (
	"fmt"
	"os"
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type MongoDiscovery struct {
	mongoUrl     string
	cluster      *Cluster
	mongoSession *mgo.Session
}

func (m *MongoDiscovery) Connect(mongoUrl string, cluster *Cluster) {
	m.mongoUrl = mongoUrl
	m.cluster = cluster

	if session, err := mgo.Dial(mongoUrl); err != nil {
		fmt.Printf("Cluster: Cannot connect to the MongoBD discovery: %s", err)
		os.Exit(1)
	} else {
		m.mongoSession = session
	}
}

func (m *MongoDiscovery) Register() {
	stop := make(chan bool)
	go m.ping(stop)
}

func (m *MongoDiscovery) ping(stop chan bool) {
	duration := time.Duration(m.cluster.PingInterval) * time.Millisecond

	for {
		session := m.mongoSession.Copy()
		collection := session.DB("").C("cluster-endpoints")

		select {
		case <-stop:
			session.Close()
			return
		default:

			selector := bson.M{
				"serviceName": m.cluster.ServiceName,
				"endpoint":    m.cluster.EndpointUrl,
			}

			updatedDoc := bson.M{
				"pingInterval": m.cluster.PingInterval,
				"endpointHash": m.cluster.EndpointHash,
				"timestamp":    time.Now(),
			}

			_, err := collection.Upsert(selector, bson.M{"$set": updatedDoc})
			if err != nil {
				fmt.Printf("Cluster: Ping failed: %s", err)
			}

			session.Close()
			time.Sleep(duration)
		}
	}
}
