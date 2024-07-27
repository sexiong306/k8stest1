package main

import (
	//"git.code.oa.com/trpc-go/trpc-go"
	pb "github.com/sexiong306/aibum-protocol/helloworld"
	"k8stest1/service"
	"trpc.group/trpc-go/trpc-go"
)

func main() {
	s := trpc.NewServer()
	pb.RegisterGreeterService(s, &service.GreeterImpl{})
	if err := s.Serve(); err != nil {
		println("server abort")
	}
}
