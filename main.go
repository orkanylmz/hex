package main

import (
	"database/sql"
	"hex/internal/adapters/grpc"
	"log"
)

func main() {
	// ..
	// ..
	db, err := sql.Open("postgres", "")
	if err != nil {
		panic(err)
	}

	grpcAdapter := grpc.NewAdapter(db)

	// Start the gRPC server
	if err := grpcAdapter.ServeGRPC(); err != nil {
		log.Fatalf("failed to serve gRPC: %v", err)
	}
}
