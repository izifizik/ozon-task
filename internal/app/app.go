package app

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
	"os"
	"os/signal"
	"ozon-task/internal/config"
	"ozon-task/internal/delivery/gRPC"
	"ozon-task/internal/domain/server"
	"ozon-task/internal/repository/cache"
	"ozon-task/internal/repository/database"
	"ozon-task/internal/service"
	pb "ozon-task/internal/service/proto"
	"syscall"
)

func Run() error {
	cfg := config.NewConfig()

	listener, err := net.Listen("tcp", cfg.Host+":"+cfg.Port)
	if err != nil {
		log.Fatalln(err.Error())
	}

	srv := grpc.NewServer()
	grpcServer := server.NewServer(srv)

	repo := database.NewRepo(cfg.DB)
	c := cache.NewCache()

	s := service.NewService(repo, c)
	err = s.CacheInit(context.Background())
	if err != nil {
		return err
	}

	handler := gRPC.NewHandler(s)

	pb.RegisterURLServiceServer(grpcServer.GrpcServer, handler)

	go gracefulShutdown([]os.Signal{syscall.SIGABRT, syscall.SIGQUIT, syscall.SIGHUP, os.Interrupt, syscall.SIGTERM}, grpcServer, listener)

	return grpcServer.GrpcServer.Serve(listener)
}

func gracefulShutdown(signals []os.Signal, closeItems ...io.Closer) {
	sign := make(chan os.Signal, 1)
	signal.Notify(sign, signals...)
	sig := <-sign

	log.Printf("Caught signal %s. Shutting down...", sig)
	for _, closer := range closeItems {
		err := closer.Close()
		if err != nil {
			fmt.Printf("failed to close %v: %v", closer, err)
		}
	}
}
