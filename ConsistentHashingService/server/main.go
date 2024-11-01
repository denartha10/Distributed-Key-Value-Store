package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"sync"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/denartha10/DistributedKeyValueStore/ConsistentHashingService/hashing"
	"github.com/denartha10/DistributedKeyValueStore/ConsistentHashingService/pb"
)

var fileMutex sync.Mutex

func fileWriter(s string) {
	fileMutex.Lock()
	defer fileMutex.Unlock()

	f, err := os.OpenFile("./nodeslog.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Fatalln("Could not open file")
	}
	f.Write([]byte(s))
}

type server struct {
	pb.UnimplementedRequestsServer
	StorageIndex hashing.AVLTree
}

type ChunkServer struct {
	IP   string
	Port int32
}

// So the function has to have the same name as the one specified in the protobuf
func (s *server) AddNode(ctx context.Context, in *pb.AddRequest) (*pb.AddResponse, error) {
	chunkServerDetails := ChunkServer{in.IP, in.Port}
	newNode := hashing.NewAVLNode(hashing.HashFunction(in.Name), chunkServerDetails)
	s.StorageIndex.Insert(newNode)
	go fileWriter(fmt.Sprintf("%s,%s,%d\n", in.Name, chunkServerDetails.IP, chunkServerDetails.Port))

	return &pb.AddResponse{StatusAdded: true}, nil
}

func main() {
	go fileWriter(fmt.Sprintf("Name,IP,Port\n"))
	// The reason we are using a TCP listener instead of the http packet is that
	// gRPC uses http2 and the http package has to support 1.1 too
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln("Failed to create listener:", err)
	}

	s := grpc.NewServer()
	reflection.Register(s)

	pb.RegisterRequestsServer(s, &server{})

	if err := s.Serve(listener); err != nil {
		log.Fatalln("failed to serve:", err)
	}
}
