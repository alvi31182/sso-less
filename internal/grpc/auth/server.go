package auth

import (
	ssov1 "github.com/alvi31182/grpc-protos"
)

type ServerAPI struct {
	ssov1.UnimplementedAuthServiceServer
}
