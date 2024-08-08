package auth

import (
	"context"
	"google.golang.org/grpc"
	auth "messenger/protobuf/generated/proto"
)

type authAPI struct {
	auth.UnimplementedUserAuthServer
}

func Register(gRPC *grpc.Server) {
	auth.RegisterUserAuthServer(gRPC, &authAPI{})
}

func (s *authAPI) Login(
	ctx context.Context,
	req *auth.LoginRequest,
) (auth.LoginResponse, error) {
	panic("no code dude")
}

func (s *authAPI) Register(
	ctx context.Context,
	req *auth.RegisterRequest,
) (auth.RegisterResponse, error) {
	panic("no code dude")
}
