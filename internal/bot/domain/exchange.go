package domain

import "time"

type Exchange interface {
	GetInfo() (*ExchangeInfo, error)
	GetOrderState(ID string) (*ExchangeOrderState, error)
}

type ExchangeOrderState struct {
}

type ExchangeInfo struct {
	Quote float32
	Epoch time.Time
}

func NewExchange() Exchange {
	return &mockExchange{}
}

type mockExchange struct {
}

func (e *mockExchange) GetInfo() (*ExchangeInfo, error) {
	return &ExchangeInfo{0.0, time.Now()}, nil
}

func (e *mockExchange) GetOrderState(ID string) (*ExchangeOrderState, error) {
	return &ExchangeOrderState{}, nil
}
