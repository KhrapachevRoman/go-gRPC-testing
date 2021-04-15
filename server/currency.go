package server

import (
	"context"
	"log"

	"github.com/KhrapachevRoman/currency/data"
	protos "github.com/KhrapachevRoman/currency/protos/currency"
)

// Currency is a gRPC server it implements the methods defined by the CurrencyServer interface
type Currency struct {
	rates *data.ExchangeRates
	log   *log.Logger
	protos.UnimplementedCurrencyServer
}

// NewCurrency creates a new Currency server
func NewCurrency(r *data.ExchangeRates, l *log.Logger) *Currency {
	return &Currency{rates: r, log: l}
}

// GetRate implements the CurrencyServer GetRate method and returns the currency exchange rate
// for the two given currencies.
func (c *Currency) GetRate(ctx context.Context, in *protos.RateRequest) (*protos.RateResponse, error) {
	c.log.Println("Handle GetRate", "base", in.GetBase(), "destination", in.GetDestination())
	rate, err := c.rates.GetRate(in.GetBase().String(), in.GetDestination().String())
	if err != nil {
		return nil, err
	}
	return &protos.RateResponse{Rate: rate}, nil
}
