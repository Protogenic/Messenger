package auth

import (
	"context"
	"google.golang.org/grpc"
	auth "messenger/protobuf/generated/proto"
)

type authAPI struct {
	auth.UnimplementedUserAuthServer
}

func Register(grpcServer *grpc.Server) {
	auth.RegisterUserAuthServer(grpcServer, &authAPI{})
}

func (a *authAPI) Login(ctx context.Context, in *auth.LoginRequest) (*auth.LoginResponse, error) {
	panic("eshkeerreee")
}
func (a *authAPI) Register(ctx context.Context, in *auth.RegisterRequest) (*auth.RegisterResponse, error) {
	panic("eshkeerreee")
}
