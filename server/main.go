package main

import (
	"context"
	"fmt"
	"grpc/compiledGo"
	"io"
	"log"
	"net"

	"google.golang.org/grpc"
)

type randomRequestServerImpl struct{}

func (r *randomRequestServerImpl) Sum(ctx context.Context, req *compiledGo.SumRequest) (*compiledGo.SumResponse, error) {
	total := req.First + req.Second
	res := &compiledGo.SumResponse{
		TotalSum: total,
	}
	return res, nil
}

func (r *randomRequestServerImpl) PrimeNumber(req *compiledGo.PrimeNumberRequest, res compiledGo.RandomeRequest_PrimeNumberServer) error {
	value := req.Val
	fmt.Println("Hello", value)
	for i := 2; int32(i) <= value; i++ {
		cnt := func(i int) int {
			cnt := 0
			for j := 1; j <= i; j++ {
				if i%j == 0 {
					cnt++
				}
			}
			return cnt
		}(i)
		if cnt == 2 {
			res.Send(&compiledGo.PrimeNumberResponse{
				Prime: int32(i),
			})
		}
	}
	return nil
}
func (r *randomRequestServerImpl) ComputeAverage(req compiledGo.RandomeRequest_ComputeAverageServer) error {
	sum := 0
	cnt := 0
	for {
		value, err := req.Recv()
		if err == io.EOF {
			return req.SendAndClose(&compiledGo.ComputeAverageResponse{
				Average: float32(sum) / float32(cnt),
			})
		} else if err != nil {
			log.Fatal(err)
		}
		sum += int(value.Val)
		cnt++
	}
}
func (r *randomRequestServerImpl) FindMaxNumber(req compiledGo.RandomeRequest_FindMaxNumberServer) error {
	var maxx int32
	for {
		request, err := req.Recv()
		if err == io.EOF {
			return nil
		}
		if maxx < request.Number {
			maxx = request.Number
			err = req.Send(&compiledGo.FindMaxResponse{
				Maximum: maxx,
			})
			if err != nil {
				fmt.Printf("Error while seding response to client %v", err)
			}
		}
	}
}
func main() {
	fmt.Println("You are Inside Server")
	list, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Internal Issue")
		return
	}
	serverImpl := grpc.NewServer()
	compiledGo.RegisterRandomeRequestServer(serverImpl, &randomRequestServerImpl{})
	if err := serverImpl.Serve(list); err != nil {
		fmt.Println("Not able to start server")
		return
	}
}
