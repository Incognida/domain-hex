package v1

import (
	"awesomeProject/app/interface/rpc/v1.0/protocol"
	"awesomeProject/app/registry"
	"awesomeProject/app/usecase"
	"google.golang.org/grpc"
)

func Apply(server *grpc.Server, ctn *registry.Container) {
	protocol.RegisterUserServiceServer(server, NewUserService(ctn.Resolve("user-usecase").(usecase.UserUsecase)))
}
