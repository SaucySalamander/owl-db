package main

import (
	"net"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	internal_grpc "github.com/SaucySalamander/owl-db/pkg/gRPC"
	"google.golang.org/grpc"

	_ "github.com/lib/pq"
)

func main() {
	if err := setup_env(); err != nil {
		panic(err)
	}
	var host = viper.GetString("host")
	var port = viper.GetString("port")

	lis, err := net.Listen("tcp", host+":"+port)
	if err != nil {
		log.Errorf("Failure to create server, %v", err)
	}

	server := grpc.NewServer()
	internal_grpc.RegisterServer(server)

	if err := server.Serve(lis); err != nil {
		log.Errorf("Failed to serve the grpc server")
	}
}

func setup_env() error {
	viper.SetConfigName("config_local")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	return viper.ReadInConfig()
}
