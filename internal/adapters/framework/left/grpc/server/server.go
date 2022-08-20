package server

import (
	"context"

	"go.hexagonal-architecture/internal/ports"
	pb2 "go.hexagonal-architecture/internal/ports/framework_left"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Adapter struct {
	api ports.ApiPort
	pb2.UnimplementedArithmeticServiceServer
}

func NewAdapter(api ports.ApiPort) *Adapter {
	return &Adapter{
		api: api,
	}
}

func (a Adapter) GetAddition(_ context.Context, req *pb2.OperationParameters) (*pb2.Answer, error) {
	ans := &pb2.Answer{}

	if req.GetX() == 0 || req.GetY() == 0 {
		return ans, status.Error(codes.InvalidArgument, "missing required")
	}

	answer, err := a.api.GetAddition(req.X, req.Y)
	if err != nil {
		return ans, status.Error(codes.Internal, "unexpected error")
	}

	ans = &pb2.Answer{
		Value: answer,
	}
	return ans, nil
}

func (a Adapter) GetSubtraction(_ context.Context, req *pb2.OperationParameters) (*pb2.Answer, error) {
	ans := &pb2.Answer{}

	if req.GetX() == 0 || req.GetY() == 0 {
		return ans, status.Error(codes.InvalidArgument, "missing required")
	}

	answer, err := a.api.GetSubtraction(req.X, req.Y)
	if err != nil {
		return ans, status.Error(codes.Internal, "unexpected error")
	}

	ans = &pb2.Answer{
		Value: answer,
	}
	return ans, nil
}

func (a Adapter) GetMultiplication(_ context.Context, req *pb2.OperationParameters) (*pb2.Answer, error) {
	ans := &pb2.Answer{}

	if req.GetX() == 0 || req.GetY() == 0 {
		return ans, status.Error(codes.InvalidArgument, "missing required")
	}

	answer, err := a.api.GetMultiplication(req.X, req.Y)
	if err != nil {
		return ans, status.Error(codes.Internal, "unexpected error")
	}

	ans = &pb2.Answer{
		Value: answer,
	}
	return ans, nil
}

func (a Adapter) GetDivision(_ context.Context, req *pb2.OperationParameters) (*pb2.Answer, error) {
	ans := &pb2.Answer{}

	if req.GetX() == 0 || req.GetY() == 0 {
		return ans, status.Error(codes.InvalidArgument, "missing required")
	}

	answer, err := a.api.GetDivision(req.X, req.Y)
	if err != nil {
		return ans, status.Error(codes.Internal, "unexpected error")
	}

	ans = &pb2.Answer{
		Value: answer,
	}
	return ans, nil
}
