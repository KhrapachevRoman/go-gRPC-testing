package main

import (
	"log"
	"net"
	"os"

	protos "github.com/KhrapachevRoman/currency/protos/currency"
	"github.com/KhrapachevRoman/currency/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {

	log := log.New(os.Stdout, "products-api", log.LstdFlags)

	gs := grpc.NewServer()
	cs := server.NewCurrency(log)
	protos.RegisterCurrencyServer(gs, cs)

	reflection.Register(gs)

	l, err := net.Listen("tcp", ":9092")
	if err != nil {
		log.Fatal("Unable to listen", "error", err)
		os.Exit(1)
	}
	gs.Serve(l)
}
