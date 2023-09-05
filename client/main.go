package main

import (
	"context"
	"fmt"
	h "inventory/proto"
	"log"

	"google.golang.org/grpc"
)

func main(){
	conn,err := grpc.Dial("localhost:2023",grpc.WithInsecure())
	if err!=nil{
		fmt.Println("Failed to connect: ",err)
	}
	defer conn.Close()
	client := h.NewInventoryServiceClient(conn)

	response, err := client.GetAllItems(context.Background(), &h.Empty{})
	if err!=nil{
		log.Fatal("Failed to call: ",err)
	}
	fmt.Println("Response: ",response)
}