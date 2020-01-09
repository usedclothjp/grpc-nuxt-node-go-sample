package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/usedclothjp/grpc-nuxt-node-go-sample/server/sumpb"
	"google.golang.org/grpc"
)

type server struct{}

func (*server) Sum(ctx context.Context, req *sumpb.SumRequest) (*sumpb.SumResponse, error) {
	fmt.Println("got access!")
	fmt.Printf("num1: %v\n", req.GetNum1())
	fmt.Printf("num2: %v\n", req.GetNum2())

	result := req.GetNum1() + req.GetNum2()

	return &sumpb.SumResponse{
		Num: result,
	}, nil
}

func main() {
	fmt.Println("sum service")
	lis, err := net.Listen("tcp", "0.0.0.0:60000")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	sumpb.RegisterSumServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to server: %v", err)
	}
}
