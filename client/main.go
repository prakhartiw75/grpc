package main

import (
	"context"
	"fmt"
	"grpc/compiledGo"
	"io"
	"log"

	"google.golang.org/grpc"
)

func ExecuteSum(client compiledGo.RandomeRequestClient) {
	first, second := 2, 3
	request := &compiledGo.SumRequest{
		First:  int32(first),
		Second: int32(second),
	}
	response, err := client.Sum(context.Background(), request)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(response.TotalSum)
}

func ExecutePrimeNumber(client compiledGo.RandomeRequestClient) {
	value := 100
	request := &compiledGo.PrimeNumberRequest{
		Val: int32(value),
	}
	responseStream, err := client.PrimeNumber(context.Background(), request)
	if err != nil {
		fmt.Printf("Error Occured=%s", err.Error())
	}
	for {
		value, err := responseStream.Recv()
		if err == io.EOF {
			fmt.Println("All inputes are read. Terminating!")
			break
		} else if err != nil {
			fmt.Println("Unknown Error Occured! Stopping Process")
		} else {

			fmt.Println("Prime No is", value.Prime)
		}
	}
}

func ExecuteCalculateAverage(client compiledGo.RandomeRequestClient) {
	requestStream, err := client.ComputeAverage(context.Background())
	if err != nil {
		log.Fatal("Not able to implement requestStream object")
	}
	allRequest := []*compiledGo.ComputeAverageRequest{
		{
			Val: 10,
		},
		{
			Val: 20,
		},
		{
			Val: 30,
		},
		{
			Val: 40,
		},
		{
			Val: 50,
		},
		{
			Val: 60,
		},
		{
			Val: 70,
		},
	}
	for _, val := range allRequest {
		fmt.Println("Request Sent=>", val.Val)
		requestStream.Send(val)
	}
	finalAns, err := requestStream.CloseAndRecv()
	if err != nil {
		fmt.Println("Error Occured while closing connection")
	}
	fmt.Println(finalAns.Average)
}

func main() {
	fmt.Println("You Are Inside client")
	channel, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		fmt.Println("Problem with establishing Connection")
		return
	}
	client := compiledGo.NewRandomeRequestClient(channel)
	//ExecuteSum(client)
	//ExecutePrimeNumber(client)
	ExecuteCalculateAverage(client)
}
