package gapi

import (
	"context"

	"github.com/lib/pq"
	db "github.com/tomoki-yamamura/simple_bank/db/sqlc"
	"github.com/tomoki-yamamura/simple_bank/pb"
	"github.com/tomoki-yamamura/simple_bank/util"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	hashedPassword, err := util.HashPassword(req.GetPassword())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "fail to hashpassword")
	}

	arg := db.CreateUserParams{
		Username: req.GetUsername(),
		HashedPassword: hashedPassword,
		FullName: req.GetFullName(),
		Email: req.GetEmail(),
	}

	user, err := server.store.CreateUser(ctx, arg)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "unique_violation":
				return nil, status.Errorf(codes.AlreadyExists, "username already exists %s", err)
			}
		}
		return nil, status.Errorf(codes.Internal, "fail to create user %s", err)
	}
	response := pb.CreateUserResponse{
		User: convertUser(user),
	}
	return &response, nil
}
