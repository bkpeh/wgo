package main

import (
	"context"
	"fmt"

	datamsg "github.com/bkpeh/wgo/proto"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func main() {
	//Create a client connection
	conn, err := grpc.Dial(":50000", grpc.WithInsecure())

	if err != nil {
		fmt.Println("Fail to dial.", err)
	}

	defer conn.Close()

	//Set the client connection to the GRPC API
	client := datamsg.NewSetTransactionInfoClient(conn)

	t := &datamsg.Transaction{
		TransactionId:   "123",
		TransactionTime: timestamppb.Now(),
		Customer: &datamsg.Person{
			Name:    "Name",
			Address: "Address",
			Phone:   123,
		},
		Bookslist: []*datamsg.Books{
			{Title: "title",
				Subject: "subject",
				Price:   1.1,
				Stream:  datamsg.StreamType_EXPRESS,
				Level:   datamsg.LevelType_SECONDARY1,
			},
		},
	}

	//Trigger the GRPC
	respond, err := client.SetTransaction(context.Background(), t)

	if err != nil {
		fmt.Println("Error in respond.", err)
	}

	fmt.Println("Respond Code.", respond.Code)
}
