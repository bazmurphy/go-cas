package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"sync"

	"github.com/bazmurphy/go-cas/keyvaluestore"

	"google.golang.org/grpc"
)

var (
	serverID = flag.Int("id", 1, "the server id")
	port     = flag.Int("port", 8080, "the TCP port the server will listen on")
)

func main() {
	flag.Parse()

	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("error: failed to listen on TCP Port 8080: %v", err)
	}

	grpcServer := grpc.NewServer()

	server := &Server{
		store: make(map[string]int64),
	}

	keyvaluestore.RegisterKeyValueStoreServiceServer(grpcServer, server)

	log.Println("starting gRPC Server...")

	err = grpcServer.Serve(listen)
	if err != nil {
		log.Fatalf("error: grpc Server failed: %v", err)
	}
}

type Server struct {
	mu    sync.RWMutex
	store map[string]int64

	keyvaluestore.UnimplementedKeyValueStoreServiceServer
}

func (s *Server) Get(ctx context.Context, req *keyvaluestore.GetRequest) (*keyvaluestore.GetResponse, error) {
	// log.Printf("Server %d | GET | Request Received: %v", *serverID, req)
	s.mu.RLock()
	defer s.mu.RUnlock()
	value, ok := s.store[req.Key]
	if !ok {
		return &keyvaluestore.GetResponse{}, nil
	}
	return &keyvaluestore.GetResponse{Value: value}, nil
}

func (s *Server) Set(ctx context.Context, req *keyvaluestore.SetRequest) (*keyvaluestore.SetResponse, error) {
	// log.Printf("Server %d | SET | Request Received: %v", *serverID, req)
	s.mu.Lock()
	defer s.mu.Unlock()
	s.store[req.Key] = req.Value
	return &keyvaluestore.SetResponse{Success: true}, nil
}

func (s *Server) Cas(ctx context.Context, req *keyvaluestore.CasRequest) (*keyvaluestore.CasResponse, error) {
	log.Printf("Server %d | CAS | Request Received: %v", *serverID, req)
	s.mu.Lock()
	defer s.mu.Unlock()
	value, ok := s.store[req.Key]
	if !ok || value != req.ExpectedValue {
		log.Printf("Server %d | CAS | Failed | Key: %s | ExpectedValue: %d | NewValue: %d", *serverID, req.Key, req.ExpectedValue, req.NewValue)
		return &keyvaluestore.CasResponse{Success: false}, nil
	}
	s.store[req.Key] = req.NewValue
	log.Printf("Server %d | CAS | Success | Key: %s | ExpectedValue: %d | NewValue: %d", *serverID, req.Key, req.ExpectedValue, req.NewValue)
	return &keyvaluestore.CasResponse{Success: true}, nil
}
