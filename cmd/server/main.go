package main

import (
	"fmt"
	"grpc/ives/pb"
	"log"
	"math/rand"
	"net"
	"time"

	"google.golang.org/grpc"
)

type streamExampleServer struct {
	pb.UnimplementedStreamExampleServer
}

func (s *streamExampleServer) SendQuotation(req *pb.QuotationRequest, stream pb.StreamExample_SendQuotationServer) error {
	for {
		randomFloat := rand.Float32()

		select {
		case <-stream.Context().Done():
			// Stream encerrada
			fmt.Println("Stream encerrada")
			return nil
		default:
			// Continua processando a stream
		}
		currentTime := time.Now()
		formattedTime := currentTime.Format(time.RFC3339)
		resp := &pb.QuotationNotify{Price: randomFloat, Message: formattedTime}
		stream.Send(resp)

		time.Sleep(time.Second)
	}
}

func (s *streamExampleServer) SendData(req *pb.SendDataRequest, stream pb.StreamExample_SendDataServer) error {
	numMessages := req.NumMessages

	for i := 1; i <= int(numMessages); i++ {
		message := fmt.Sprintf("Message %d", i)
		resp := &pb.SendDataResponse{Message: message}
		if err := stream.Send(resp); err != nil {
			return err
		}
		time.Sleep(time.Second)
	}

	return nil
}

func main() {
	port := 50051
	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	server := grpc.NewServer()
	pb.RegisterStreamExampleServer(server, &streamExampleServer{})

	log.Printf("Server listening on port %d", port)
	if err := server.Serve(listen); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
