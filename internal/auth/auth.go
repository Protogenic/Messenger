package auth

import auth "protos/generated/proto"

type authAPI struct {
	auth.UnimplementedUserAuthServer
}
