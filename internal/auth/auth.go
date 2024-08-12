package auth

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	authpb "messenger/protobuf/generated/proto"
)

type Auth interface {
	Login(ctx context.Context, request *authpb.LoginRequest) (token string, err error)
	Register(ctx context.Context, request *authpb.RegisterRequest) (userID int64, err error)
}

type authAPI struct {
	authpb.UnimplementedUserAuthServer
	auth Auth
}

func Register(grpcServer *grpc.Server, auth Auth) {
	authpb.RegisterUserAuthServer(grpcServer, &authAPI{auth: auth})
}

func (a *authAPI) Login(ctx context.Context, req *authpb.LoginRequest) (*authpb.LoginResponse, error) {
	if req.GetUsername() == "" {
		return nil, status.Errorf(codes.InvalidArgument, "username is empty")
	}
	if req.GetPassword() == "" {
		return nil, status.Errorf(codes.InvalidArgument, "password is empty")
	}

	token, err := a.auth.Login(ctx, req)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "polomka error")
	}

	return &authpb.LoginResponse{
		Token: token,
	}, nil
}

func (a *authAPI) Register(ctx context.Context, req *authpb.RegisterRequest) (*authpb.RegisterResponse, error) {
	if req.GetUsername() == "" {
		return nil, status.Errorf(codes.InvalidArgument, "username is empty")
	}
	if req.GetPassword() == "" {
		return nil, status.Errorf(codes.InvalidArgument, "password is empty")
	}
	userID, err := a.auth.Register(ctx, req)
	if err != nil {
		// TODO: ... already exist
		return nil, status.Errorf(codes.Internal, "polomka error")
	}
	return &authpb.RegisterResponse{
		UserId: userID,
	}, nil
}
