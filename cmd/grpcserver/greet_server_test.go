package main_test

import (
	"fmt"
	"go-specs-greet/adapters"
	"go-specs-greet/adapters/grpcserver"
	"go-specs-greet/specifications"
	"testing"
)

func TestGreeterServer(t *testing.T) {
	var (
		port           = "50051"
		dockerfilePath = "./cmd/grpcserver/Dockerfile"
		driver         = grpcserver.Driver{Addr: fmt.Sprintf("localhost:%s", port)}
	)

	adapters.StartDockerServer(t, port, dockerfilePath)
	specifications.GreetSpecification(t, &driver)
}