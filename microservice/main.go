package microservice

import (
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
	"log"
	"net"
)

type MicroApp struct {
	server          *grpc.Server
	grpcPort        string
	registerService func(*grpc.Server)
}

func (app *MicroApp) register() {
	log.Println("Registering gRPC server..." + app.grpcPort)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", app.grpcPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	app.registerService(app.server)

	log.Printf("gRPC Server started on port %s", app.grpcPort)

	if err := app.server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (app *MicroApp) checkHealth() {
	healthServer := health.NewServer()
	healthServer.SetServingStatus("", healthpb.HealthCheckResponse_SERVING)
	healthpb.RegisterHealthServer(app.server, healthServer)
}
