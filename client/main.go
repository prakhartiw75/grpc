package main

import (
	"fmt"
	"grpc/compiledGo"
)

func main() {
	fmt.Println("inside client")
	dum := compiledGo.Dummy{
		Name: "Prakhar Tiwari",
	}
	fmt.Println(dum.GetName())
}
