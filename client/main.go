package main

import (
	"context"
	"flag"
	"log"

	"github.com/bazmurphy/go-cas/keyvaluestore"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	serverAddress = flag.String("server", "", "the server address to connect to")
	clientID      = flag.Int("id", 1, "the client id")
	requests      = flag.Int("requests", 1, "the number of CAS requests to make")
)

func main() {
	flag.Parse()

	client, err := NewClient(*serverAddress)
	if err != nil {
		log.Fatalf("error: failed to connect to the gRPC Server: %v", err)
	}

	if *clientID == 1 {
		// setResponse, err := client.Set("baz", 0)
		// if err != nil {
		// 	log.Printf("Client %d | SET | Error: %v", *clientID, err)
		// } else {
		// 	log.Printf("Client %d | SET | Response: %v", *clientID, setResponse)
		// }
		// getResponse, err := client.Get("baz")
		// if err != nil {
		// 	log.Printf("Client %d | GET | Error: %v", *clientID, err)
		// } else {
		// 	log.Printf("Client %d | GET | Response: %v", *clientID, getResponse)
		// }

		// initially set a value on the key "baz"
		client.Set("baz", 0)
		// client.Get("baz")
	}

	// make a bunch of CAS requests to the key "baz"
	for i := 0; i <= *requests; i++ {
		casResponse, err := client.Cas("baz", int64(i), int64(i+1))
		if err != nil {
			log.Printf("Client %d | CAS | Error: %v", *clientID, err)
		} else {
			log.Printf("Client %d | CAS | Response: %v", *clientID, casResponse)
		}
	}
}

type KeyValueStore interface {
	Get(key string) (*keyvaluestore.GetResponse, error)
	Set(key string, value int64) (*keyvaluestore.SetResponse, error)
	Cas(key string, expectedValue, newValue int64) (*keyvaluestore.CasResponse, error)
}

type Client struct {
	client keyvaluestore.KeyValueStoreServiceClient
}

func NewClient(address string) (KeyValueStore, error) {
	connection, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	client := keyvaluestore.NewKeyValueStoreServiceClient(connection)
	return &Client{client: client}, nil
}

func (c *Client) Get(key string) (*keyvaluestore.GetResponse, error) {
	getRequest := &keyvaluestore.GetRequest{
		Key: key,
	}
	getResponse, err := c.client.Get(context.Background(), getRequest)
	if err != nil {
		return nil, err
	}
	return getResponse, nil
}

func (c *Client) Set(key string, value int64) (*keyvaluestore.SetResponse, error) {
	setRequest := &keyvaluestore.SetRequest{
		Key:   key,
		Value: value,
	}
	setResponse, err := c.client.Set(context.Background(), setRequest)
	if err != nil {
		return nil, err
	}
	return setResponse, nil
}

func (c *Client) Cas(key string, expectedValue, newValue int64) (*keyvaluestore.CasResponse, error) {
	casRequest := &keyvaluestore.CasRequest{
		Key:           key,
		ExpectedValue: expectedValue,
		NewValue:      newValue,
	}
	casResponse, err := c.client.Cas(context.Background(), casRequest)
	if err != nil {
		return nil, err
	}
	return casResponse, nil
}
