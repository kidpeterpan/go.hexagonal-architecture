package main

import (
	"log"
	"net"

	"go.hexagonal-architecture/internal/adapters/app/api"
	"go.hexagonal-architecture/internal/adapters/core/arithmetic"
	"go.hexagonal-architecture/internal/adapters/framework/left/grpc/server"
	"go.hexagonal-architecture/internal/adapters/framework/right/db"
	"go.hexagonal-architecture/internal/ports"
	"go.hexagonal-architecture/internal/ports/framework_left"
	"google.golang.org/grpc"
)

func main() {

	// ports
	var core ports.ArithmeticPort
	var database ports.DbPort
	var app ports.ApiPort
	var arithmeticServiceServer framework_left.ArithmeticServiceServer

	// plugin adapter
	core = arithmetic.NewAdapter()
	database = db.NewAdapter()
	app = api.NewAdapter(core, database)
	arithmeticServiceServer = server.NewAdapter(app)

	defer func() {
		if err := database.CloseDbConnection(); err != nil {
			log.Fatalf("CloseDbConnection error: %v", err)
		}
	}()

	listen, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen on port 50051: %v", err)
	}

	grpcServer := grpc.NewServer()
	framework_left.RegisterArithmeticServiceServer(grpcServer, arithmeticServiceServer)

	log.Println("listening on port 50051")

	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("failed to serve gRPC server over on port 9000: %v", err)
	}

}
