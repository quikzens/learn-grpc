package main

import (
	pb "bookshop/client/pb/inventory"
	"context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))

	// Code removed for brevity

	client := pb.NewInventoryClient(conn)

	// Note how we are calling the GetBookList method on the server
	// This is available to us through the auto-generated code
	bookList, err := client.GetBookList(context.Background(), &pb.GetBookListRequest{})
	if err != nil {
		log.Println(err)
	}

	log.Printf("book list: %v", bookList)
}
