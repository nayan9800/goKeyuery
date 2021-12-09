package server

import (
	"context"
	"log"
	"net"

	"github.com/nayan9800/goKeyuery/data"
	pb "github.com/nayan9800/goKeyuery/protoc"
	"google.golang.org/grpc"
)

//struct KeyvalueServer have
//UnimplementedKeyvalueServer
type KeyvalueServer struct {
	pb.UnimplementedKeyvalueServer
}

//start server function
func StartServer(address string) {
	//tcp listner for given port
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal(err.Error())
	}
	//create grpc server object
	server := grpc.NewServer()

	//register the ketValueServer with grpc
	//server
	pb.RegisterKeyvalueServer(server, &KeyvalueServer{})
	log.Println("Listeninig on :5000")
	//start server
	if err := server.Serve(lis); err != nil {
		log.Fatal(err.Error())
	}
}

/*
	StoreKeyValue stores the key value pair in
	database
*/
func (s *KeyvalueServer) StoreKeyValue(ctx context.Context, kv *pb.KeyValue) (*pb.Value, error) {
	data.SaveKeyValue(kv)
	return kv.Value, nil
}

/*
	Getvalue gets the value of given key
	from database
*/
func (s *KeyvalueServer) GetValue(ctx context.Context, k *pb.Key) (*pb.Value, error) {
	v := data.GetValue(k)
	return v, nil
}
