package main

import (
	"log"
	"net"

	service "echo_v1"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var (
	launcher *grpc.Server
	gservice service.EchoServer
)

func shutdown() {
	if launcher == nil {
		return
	}
	launcher.GracefulStop()
}

func serve1(address string) {
	ln, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalln(err)
	}
	defer ln.Close()

	launcher = grpc.NewServer()
	service.RegisterEchoServer(launcher, gservice)

	if err := launcher.Serve(ln); err != nil {
		log.Fatalln(err)
	}
}

func serve2(address, keypath, certpath string) {
	ln, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalln(err)
	}
	defer ln.Close()

	creds, err := credentials.NewServerTLSFromFile(certpath, keypath)
	if err != nil {
		log.Fatalln(err)
	}
	launcher := grpc.NewServer(grpc.Creds(creds))
	service.RegisterEchoServer(launcher, gservice)

	if err := launcher.Serve(ln); err != nil {
		log.Fatalln(err)
	}
}
