package rpc

import (
	"awesomeProject/app/interface/rpc/v1.0"
	"awesomeProject/app/registry"
	"google.golang.org/grpc"
)

func Apply(server *grpc.Server, ctn *registry.Container) {
	v1.Apply(server, ctn)
}
