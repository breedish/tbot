package domain

import "time"

type Exchange interface {
	GetInfo() ExchangeInfo
	GetOrderState(ID string) ExchangeOrderState
}

type ExchangeOrderState interface {
}

type ExchangeInfo struct {
	Quote float32
	Epoch time.Time
}
