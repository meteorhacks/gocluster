package gocluster

import (
	"fmt"
  "time"
  "gopkg.in/mgo.v2"
  "gopkg.in/mgo.v2/bson"
  "os"
)

type Payload struct {
  Timestamp time.Time `bson: "timestamp"`
  PingInterval int64 `bson: "pingInterval"`
  Endpoint string `bson: "endpoint"`
  EndpointHash string `bson: "endpointHash"`
  ServiceName string  `bson: "serviceName"`
}

type Cluster struct {
	MongoUrl string
  ServiceName string
  EndpointUrl string
  PingInterval int64
  EndpointHash string

  mongoSession *mgo.Session
}

func (c *Cluster) Register(serviceName string, endpointUrl string) {
  c.ServiceName = serviceName
  c.EndpointUrl = endpointUrl
  c.PingInterval = 1000 * 1
  c.EndpointHash = ToSha1(endpointUrl)

	fmt.Println("Cluster: Service registered as: ", serviceName)
  fmt.Println("Cluster: Endpoint URL is: ", endpointUrl)

  stopPing := make(chan bool, 1)
  go c.ping(stopPing)

  fmt.Println(<- stopPing)
}

func (c *Cluster) ping(stop chan bool){
  duration := time.Duration(c.PingInterval) * time.Millisecond

  for {
    select {
      case <- stop:
        return
      default:
        payload := c.getPayload()
        fmt.Println(payload)

        time.Sleep(duration)
    }
  }
}

func (c *Cluster) getPayload() Payload {
  payload := Payload{}
  payload.Timestamp = time.Now()
  payload.PingInterval = c.PingInterval
  payload.Endpoint = c.EndpointUrl
  payload.EndpointHash = c.EndpointHash

  return payload
}

func Connect(mongoUrl string) *Cluster {
	cluster := &Cluster{MongoUrl: mongoUrl}

  session, err := mgo.Dial("mongodb://localhost/discovery")
  if err != nil {
    fmt.Printf("Cluster: Cannot connect to the MongoBD discovery: %s", err)
    os.Exit(1)
  } else {
    cluster.mongoSession = session
  }

  collection := session.DB("discovery").C("cluster-endpoints")

  var doc Payload
  collection.Find(bson.M{}).One(&doc)

  fmt.Println("Here's the payload: ", doc)

	return cluster
}
