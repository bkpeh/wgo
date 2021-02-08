package main

import (
	"context"
	"fmt"
	"net"

	datamsg "github.com/bkpeh/wgo/proto"
	"google.golang.org/grpc"
)

type transaction struct{}

func (trans transaction) SetTransaction(ctx context.Context, t *datamsg.Transaction) (*datamsg.TransactionReply, error) {
	tr := &datamsg.TransactionReply{
		Code: datamsg.TransactionCode_TRANSACTIONSUCCESS,
	}

	return tr, nil
}

func main() {

	listener, err := net.Listen("tcp", "50000")

	if err != nil {
		PrintLog(logerr, err)
	}

	rpcsvr := grpc.NewServer()

	fmt.Println("Running Main")
}

func httpsvr() {

}
