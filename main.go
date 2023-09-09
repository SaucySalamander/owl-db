package main

import (
	"net"

	"github.com/rs/zerolog/log"
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
		log.Fatal().Err(err)
	}

	server := grpc.NewServer()
	internal_grpc.RegisterServer(server)

	if err := server.Serve(lis); err != nil {
		log.Fatal().Err(err)
	}
}

func setup_env() error {
	viper.SetConfigName("config_local")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	return viper.ReadInConfig()
}
