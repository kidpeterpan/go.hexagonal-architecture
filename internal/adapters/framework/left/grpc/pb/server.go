package pb

import (
	"context"

	"go.hexagonal-architecture/internal/ports"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Adapter struct {
	api ports.ApiPort
}

func NewAdapter(api ports.ApiPort) *Adapter {
	return &Adapter{
		api: api,
	}
}

func (a Adapter) GetAddition(ctx context.Context, req *OperationParameters) (*Answer, error) {
	ans := &Answer{}

	if req.GetX() == 0 || req.GetY() == 0 {
		return ans, status.Error(codes.InvalidArgument, "missing required")
	}

	answer, err := a.api.GetAddition(req.X, req.Y)
	if err != nil {
		return ans, status.Error(codes.Internal, "unexpected error")
	}

	ans = &Answer{
		Value: answer,
	}
	return ans, nil
}

func (a Adapter) GetSubtraction(ctx context.Context, req *OperationParameters) (*Answer, error) {
	ans := &Answer{}

	if req.GetX() == 0 || req.GetY() == 0 {
		return ans, status.Error(codes.InvalidArgument, "missing required")
	}

	answer, err := a.api.GetSubtraction(req.X, req.Y)
	if err != nil {
		return ans, status.Error(codes.Internal, "unexpected error")
	}

	ans = &Answer{
		Value: answer,
	}
	return ans, nil
}

func (a Adapter) GetMultiplication(ctx context.Context, req *OperationParameters) (*Answer, error) {
	ans := &Answer{}

	if req.GetX() == 0 || req.GetY() == 0 {
		return ans, status.Error(codes.InvalidArgument, "missing required")
	}

	answer, err := a.api.GetMultiplication(req.X, req.Y)
	if err != nil {
		return ans, status.Error(codes.Internal, "unexpected error")
	}

	ans = &Answer{
		Value: answer,
	}
	return ans, nil
}

func (a Adapter) GetDivision(ctx context.Context, req *OperationParameters) (*Answer, error) {
	ans := &Answer{}

	if req.GetX() == 0 || req.GetY() == 0 {
		return ans, status.Error(codes.InvalidArgument, "missing required")
	}

	answer, err := a.api.GetDivision(req.X, req.Y)
	if err != nil {
		return ans, status.Error(codes.Internal, "unexpected error")
	}

	ans = &Answer{
		Value: answer,
	}
	return ans, nil
}

func (a Adapter) mustEmbedUnimplementedArithmeticServiceServer() {}
