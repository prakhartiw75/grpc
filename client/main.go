package main

import (
	"context"
	"fmt"
	"grpc/compiledGo"
	"io"
	"log"
	"strings"
	"time"

	"google.golang.org/grpc"
)

func ExecuteSum(client compiledGo.RandomeRequestClient) {
	var first, second int32 = 2, 3
	request := &compiledGo.SumRequest{
		First:  first,
		Second: second,
	}
	response, err := client.Sum(context.Background(), request)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(response.TotalSum)
}

func ExecutePrimeNumber(client compiledGo.RandomeRequestClient) {
	var value int32 = 100
	request := &compiledGo.PrimeNumberRequest{
		Val: value,
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

func ExecuteMaximumNumber(client compiledGo.RandomeRequestClient) {
	requestStream, err := client.FindMaxNumber(context.Background())
	if err != nil {
		fmt.Printf("Error while creating stream object %v", err)
	}
	allRequest := []*compiledGo.FindMaxRequest{
		{
			Number: 10,
		},
		{
			Number: 20,
		},
		{
			Number: 15,
		},
		{
			Number: 29,
		},
		{
			Number: 12,
		},
		{
			Number: 48,
		},
		{
			Number: 7,
		},
		{
			Number: 100,
		},
	}
	channel := make(chan int)
	go func(allRequest []*compiledGo.FindMaxRequest) {
		for _, val := range allRequest {
			requestStream.Send(val)
		}
		requestStream.CloseSend()
	}(allRequest)
	go func() {
		for {
			response, err := requestStream.Recv()
			if err == io.EOF {
				fmt.Println("Server is CLOSED now")
				close(channel)
				return
			}
			fmt.Println("Maximum Number is=>", response.Maximum)
		}
	}()
	<-channel
}

func main() {
	fmt.Println("You Are Inside client")
	channel, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		fmt.Println("Problem with establishing Connection")
		return
	}
	client := compiledGo.NewRandomeRequestClient(channel)
	fmt.Println("Calling Sum Service now. It will print sum of two numbers")
	ExecuteSum(client)
	time.Sleep(2 * time.Second)
	fmt.Println(strings.Repeat("#", 40))
	fmt.Println("Calling Prime Numer Service now. It will print all Prime Numbers lesser than or equal to given Number")
	ExecutePrimeNumber(client)
	time.Sleep(2 * time.Second)
	fmt.Println(strings.Repeat("#", 40))
	fmt.Println("Calling Average Number Service now. It will input stream of values and calculates average and return to client")
	ExecuteCalculateAverage(client)
	time.Sleep(2 * time.Second)
	fmt.Println(strings.Repeat("#", 40))
	fmt.Println("Calling FindMaximum Service now. It will input stream of values and whenever maximum value is acheived, it will return to client immediately")
	ExecuteMaximumNumber(client)
}
