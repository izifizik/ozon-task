package gRPC

import (
	"context"
	"ozon-task/internal/service"
	pb "ozon-task/internal/service/proto"
)

type Handler struct {
	s service.Service
	pb.UnimplementedURLServiceServer
}

func NewHandler(s service.Service) *Handler {
	return &Handler{s: s}
}

func (h *Handler) Create(ctx context.Context, url *pb.URL) (*pb.URL, error) {
	short, err := h.s.Create(ctx, url.GetBody())
	if err != nil {
		return nil, err
	}

	return &pb.URL{Body: short}, nil
}

func (h *Handler) Get(ctx context.Context, url *pb.URL) (*pb.URL, error) {
	u, err := h.s.Get(ctx, url.GetBody())
	if err != nil {
		return nil, err
	}

	return &pb.URL{Body: u}, nil
}
