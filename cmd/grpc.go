package cmd

import (
	"ewallet-framework-1/helpers"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"log"
	"net"
)

func ServeGRPC() {
	lis, err := net.Listen("tcp", ":"+helpers.GetEnv("GRPC_PORT", "7000"))
	if err != nil {
		log.Fatal("failed to listen grpc port: ", err)
	}

	s := grpc.NewServer()

	// list method
	// pb.ExampleMethod(s, &grpc....)

	logrus.Info("start listening grpc on port:" + helpers.GetEnv("GRPC_PORT", "7000"))
	if err := s.Serve(lis); err != nil {
		log.Fatal("failed to serve grpc port: ", err)
	}
}
