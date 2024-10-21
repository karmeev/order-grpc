package main

import "log"

func main() {
	grpcServer := NewGRPCServer(":9000")
	err := grpcServer.Run()
	if err != nil {
		log.Fatalf("failed to start: %v", err)
	}
}
