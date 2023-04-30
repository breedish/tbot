package domain

import (
	"time"
)

type Bot interface {
	GetState() BotState
	Init() error
	Run() error
}

type bot struct {
	ID       string
	State    BotState
	Settings Settings
	Grid     Grid
	History  []HistoryEntry
}

type Settings struct {
	BaseCurrency        Currency // Currency to purchase for future sale, e.g. BTC. Bot purchases BaseCurrency to sell it for QuoteCurrency in the future.
	QuoteCurrency       Currency // Currency to maximize profit, e.g. USDT
	QuoteCurrencyAmount float32
	LowPriceThreshold   int
	HighPriceThreshold  int
	Steps               int
	ProfitPercentage    float32
	FeePercentage       float32
}

type BotState int

const (
	Uninitialized BotState = iota
	Initialized
	Running
	Paused
	Stopped
)

type Grid struct {
	Quote float32 // Exchange rate from BaseCurrency to QuoteCurrency, e.g. BTC/USDT = 29230
	Lines []GridLine
	Epoch time.Time
}

type GridLine struct {
	ID       string
	Sell     Order // Sell BaseCurrency(BTC) to QuoteCurrency(USDT)
	Purchase Order // Purchase BaseCurrency(BTC) from QuoteCurrency(USDT)
}

type Order struct {
	ID                  string
	ExternalId          string
	Type                OrderType
	State               OrderState
	BaseCurrencyAmount  float32
	QuoteCurrencyAmount float32
	Quote               float32
}

type OrderType int

const (
	Sell OrderType = iota
	Purchase
)

type OrderState int

const (
	Inactive OrderState = iota
	Active
	Completed
)

type HistoryEntry struct {
	Epoch time.Time
	State BotState
	Quote float32
	Grid  Grid
}

func NewBroker(settings Settings) Bot {
	return &bot{
		State:    Uninitialized,
		Settings: settings,
	}
}

func (g *bot) Init() error {
	panic("Not implemented.")
}

func (g *bot) Run() error {
	panic("Not implemented.")
}

func (g *bot) GetState() BotState {
	return g.State
}
