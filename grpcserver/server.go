package grpcserver

import (
	"context"
	"github.com/go-chassis/go-chassis"
	"github.com/go-chassis/go-chassis/core/server"
	_ "github.com/go-chassis/go-chassis/server/grpc"
	"github.com/tomlee0201/chassisdemo/protobuf"
)

type HelloServer struct {
}

func (s *HelloServer) SayHello(ctx context.Context, in *protobuf.HelloRequest) (*protobuf.HelloReply, error) {
	return &protobuf.HelloReply{Message: "Hello  " + in.Name + " from rpc server"}, nil
}

func Run()  {
	chassis.RegisterSchema("grpc", &HelloServer{}, server.WithGRPCServiceDesc(&protobuf.Greeter_serviceDesc))
}
