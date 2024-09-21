package auth

import (
	"context"

	pb "github.com/cstati/auth/pkg/auth"
)

func (s *Service) Ping(context.Context, *pb.PingRequest) (*pb.PingResponse, error) {
	return &pb.PingResponse{}, nil
}
