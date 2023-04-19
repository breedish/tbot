package app

import "github.com/breedish/tbot/internal/bot/app/command"

type ApplicationUseCases struct {
	Commands Commands
	Queries  Queries
}

type Commands struct {
	BillingAppleCommand command.BillingAppleCommandHandler
}

type Queries struct {
}
