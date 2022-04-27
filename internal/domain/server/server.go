package server

import "google.golang.org/grpc"

type Server struct {
	GrpcServer *grpc.Server
}

//NewServer - for implement io.Closer
func NewServer(grpcServer *grpc.Server) *Server {
	return &Server{GrpcServer: grpcServer}
}

func (s *Server) Close() error {
	s.GrpcServer.GracefulStop()
	return nil
}
