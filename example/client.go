package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "github.com/nayan9800/goKeyuery/protoc"
	"google.golang.org/grpc"
)

func main() {

	StoreKeyValue()
	time.Sleep(2 * time.Second)
	getValue()
}

func getValue() {

	//create connection with the grpc server
	conn, err := grpc.Dial(":5000", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Println(err.Error())

	}
	defer conn.Close()
	//create the object of the Key
	k := pb.Key{V: "foo"}
	//creates the new client from connection
	client := pb.NewKeyvalueClient(conn)
	//creates the context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	//use gRPC GetValue() to get value of the
	resp, err := client.GetValue(ctx, &k)
	if err != nil {
		log.Println(err.Error())
	}
	fmt.Println(resp)
}

func StoreKeyValue() {
	//create connection with the grpc server
	conn, err := grpc.Dial(":5000", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Println(err.Error())

	}
	defer conn.Close()
	//create the keyValue object
	kv := pb.KeyValue{Key: &pb.Key{V: "apikey1"}, Value: &pb.Value{V: "jbsljbLJBFLJBAFLALJFBLJB"}}
	//creates the new client from connection
	client := pb.NewKeyvalueClient(conn)
	//creates the context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	//use gRPC StoreKeyValue() to store key value
	resp, err := client.StoreKeyValue(ctx, &kv)
	if err != nil {
		log.Println(err.Error())
	}
	fmt.Println(resp)
}
