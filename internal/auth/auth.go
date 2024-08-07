package auth

import authUser "protos/generated/proto"

type AuthApi struct {
	authUser.UnimplementedUserAuthServer
}
