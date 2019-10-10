package client

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "Juno/libraGrpc/proto"
	"google.golang.org/grpc"
)

const (
	address = "localhost:50001"
)

func ExDDLClient(order *pb.LibraAuditOrder) {

	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewJunoClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.OrderDDLExec(ctx, order)
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	fmt.Println(r.Message)
}

func ExDMLClient(order *pb.LibraAuditOrder) {

	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewJunoClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.OrderDMLExec(ctx, order)
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	fmt.Println(r.Message)
}


func TsClient(order *pb.LibraAuditOrder) {

	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewJunoClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.OrderDeal(ctx, order)
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	for _, i := range r.Record {
		fmt.Println(i.SQL)   // SQL语句
		fmt.Println(i.Status) //审核状态
		fmt.Println(i.Level)  // 错误等级
		fmt.Println(i.Error) // 错误信息
		fmt.Println(i.AffectRows) //影响行数
	}
}



