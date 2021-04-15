package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/KhrapachevRoman/currency/data"
	protos "github.com/KhrapachevRoman/currency/protos/currency"
	"github.com/KhrapachevRoman/currency/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {

	log := log.New(os.Stdout, "currency-api", log.LstdFlags)

	// generate a new ExchangeRates
	rates, err := data.NewRates(log)
	if err != nil {
		log.Fatal("Unable to generate rates, error", err)
		os.Exit(1)
	}
	// create a new gRPC server, use WithInsecure to allow http connections
	gs := grpc.NewServer()

	// create an instanse of the Currency server
	cs := server.NewCurrency(rates, log)

	// register the currency server
	protos.RegisterCurrencyServer(gs, cs)

	// register the reflection service which allows clients to determine the methods
	// for this gRPC service
	reflection.Register(gs)

	// create a TCP socket for inbound server connections
	l, err := net.Listen("tcp", fmt.Sprintf(":%d", 9092))
	if err != nil {
		log.Fatal("Unable to listen", "error", err)
		os.Exit(1)
	}
	// listen for requests
	gs.Serve(l)
}
