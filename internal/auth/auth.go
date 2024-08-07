package auth

import auth "messenger/protobuf/generated/proto"

type AuthAPI struct {
	auth.UnimplementedUserAuthServer
}
