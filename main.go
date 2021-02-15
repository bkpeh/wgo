package main

import (
	"context"
	"fmt"
	"net"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	datamsg "github.com/bkpeh/wgo/proto"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type transaction struct {
	datamsg.UnimplementedSetTransactionInfoServer
}

func (trans transaction) SetTransaction(ctx context.Context, t *datamsg.Transaction) (*datamsg.TransactionReply, error) {

	customResolver := aws.EndpointResolverFunc(func(service, region string) (aws.Endpoint, error) {
		if service == dynamodb.ServiceID && region == "us-west-2" {
			return aws.Endpoint{
				PartitionID:   "aws",
				URL:           "http://localhost:8000",
				SigningRegion: "us-west-2",
			}, nil
		}
		// returning EndpointNotFoundError will allow the service to fallback to it's default resolution
		return aws.Endpoint{}, &aws.EndpointNotFoundError{}
	})

	cfg, _ := config.LoadDefaultConfig(context.TODO(), config.WithRegion("us-west-2"), config.WithEndpointResolver(customResolver))
	svc := dynamodb.NewFromConfig(cfg)

	settransactiontoDB(svc)

	tr := &datamsg.TransactionReply{
		Code: datamsg.TransactionCode_TRANSACTIONSUCCESS,
	}

	fmt.Println("SetTransaction CB")

	return tr, nil
}

func main() {
	startserver()
}

func startserver() {
	//Start listener on TCP port 50000
	listener, err := net.Listen("tcp", ":50000")

	if err != nil {
		PrintErr("Cannot listen to port 50000.", err)
	}

	//New GRPC server
	rpcsvr := grpc.NewServer()

	//Register the implmentation to the server
	datamsg.RegisterSetTransactionInfoServer(rpcsvr, &transaction{})

	//Start serving GRPC on listener
	if err := rpcsvr.Serve(listener); err != nil {
		PrintErr("Cannot Serve GRPC.", err)
	}

	PrintInfo("Server Running...", nil)
}

func settransactiontoDB(conn *dynamodb.Client) {

	t1 := datamsg.Transaction{
		TransactionId:   "20210001",
		TransactionTime: timestamppb.Now(),
		Customer: &datamsg.Person{
			Name:    "Peter",
			Address: "Singapore",
			Phone:   1234567,
		},
		Bookslist: []*datamsg.Books{
			{
				Title:   "Speak English",
				Subject: "English",
				Price:   10.50,
				Stream:  datamsg.StreamType_EXPRESS,
				Level:   datamsg.LevelType_SECONDARY1,
			},
			{
				Title:   "Calculate",
				Subject: "Maths",
				Price:   15.80,
				Stream:  datamsg.StreamType_EXPRESS,
				Level:   datamsg.LevelType_SECONDARY1,
			},
		},
	}

	item, err := attributevalue.MarshalMap(&t1)

	if err != nil {
		PrintErr("Cannot MarshalMap", err)
	}

	_, err = conn.PutItem(context.TODO(), &dynamodb.PutItemInput{
		TableName: aws.String("TransactionTable"),
		Item:      item,
	})

	if err != nil {
		PrintErr("Cannot PutItem", err)
	}
}
