package main

import (
	"context"
	"net"

	datamsg "github.com/bkpeh/wgo/proto"
	"google.golang.org/grpc"
)

type transaction struct {
	datamsg.UnimplementedSetTransactionInfoServer
}

func (trans transaction) SetTransaction(ctx context.Context, t *datamsg.Transaction) (*datamsg.TransactionReply, error) {
	tr := &datamsg.TransactionReply{
		Code: datamsg.TransactionCode_TRANSACTIONSUCCESS,
	}

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
