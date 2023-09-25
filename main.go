package main

import (
	"context"
	"net"
	"os"
	"os/signal"

	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"

	"github.com/SaucySalamander/owl-db/cmd"
	"github.com/SaucySalamander/owl-db/internal/config"
	internal_grpc "github.com/SaucySalamander/owl-db/pkg/gRPC"
	"google.golang.org/grpc"

	_ "github.com/lib/pq"
)

func main() {
	cmd.Execute()

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()
	shutdown, err := config.InitProvider()
	if err != nil {
		log.Fatal().Err(err).Send()
	}
	defer func() {
		if err := shutdown(ctx); err != nil {
			log.Fatal().Err(err).Msg("failed to shutdown TracerProvider")
		}
	}()

	var host = viper.GetString("server.host")
	var port = viper.GetString("server.port")

	lis, err := net.Listen("tcp", host+":"+port)
	if err != nil {
		log.Fatal().Err(err).Send()
	}

	server := grpc.NewServer()
	internal_grpc.RegisterServer(server)

	log.Info().Msgf("Started server on port: %s", port)

	if err := server.Serve(lis); err != nil {
		log.Fatal().Err(err).Send()
	}
}
