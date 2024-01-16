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
	GrpcServer      *grpc.Server
	GrpcPort        string
	RegisterService func(*grpc.Server)
}

func (app *MicroApp) Register() {
	log.Println("Registering gRPC server..." + app.GrpcPort)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", app.GrpcPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	app.RegisterService(app.GrpcServer)

	log.Printf("gRPC GrpcServer started on port %s", app.GrpcPort)

	if err := app.GrpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (app *MicroApp) CheckHealth() {
	healthServer := health.NewServer()
	healthServer.SetServingStatus("", healthpb.HealthCheckResponse_SERVING)
	healthpb.RegisterHealthServer(app.GrpcServer, healthServer)
}
