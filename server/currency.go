package server

import (
	"context"
	"log"

	protos "github.com/KhrapachevRoman/currency/protos/currency"
)

// Currency is a gRPC server it implements the methods defined by the CurrencyServer interface
type Currency struct {
	l *log.Logger
	protos.UnimplementedCurrencyServer
}

// NewCurrency creates a new Currency server
func NewCurrency(log *log.Logger) *Currency {
	return &Currency{l: log}
}

// GetRate implements the CurrencyServer GetRate method and returns the currency exchange rate
// for the two given currencies.
func (c *Currency) GetRate(ctx context.Context, in *protos.RateRequest) (*protos.RateResponse, error) {
	c.l.Println("Handle GetRate", "base", in.GetBase(), "destination", in.GetDestination())

	return &protos.RateResponse{Rate: 0.5}, nil
}
