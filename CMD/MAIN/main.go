package main

import (
	"net"

	GrpcService "github.com/FACELESS-GOD/tt.git/Package/GRPCService"
	tt_git "github.com/FACELESS-GOD/tt.git/Package/ProtoSystemGenerated/PROTOBUF"
	"google.golang.org/grpc"
)

func main() {

	lis, err := net.Listen("tcp", ":8001")

	if err != nil {
		panic(err)
	}

	grpcServer := grpc.NewServer()

	tt_git.RegisterTestGRPCServiceServer(grpcServer, &GrpcService.GrpcService{})

	if err != grpcServer.Serve(lis) && err != nil {
		panic(err)
	} 
}
