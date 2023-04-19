package command

import (
	"context"

	"github.com/breedish/tbot/internal/bot/services"
	"github.com/breedish/tbot/internal/common/decorator"
	"github.com/sirupsen/logrus"
)

type BillingAppleEvent struct {
	payload []byte
}

type BillingAppleCommandHandler decorator.CommandHandler[BillingAppleEvent]

type billingAppleCommandHandler struct {
	billingService services.BillingService
}

func (h billingAppleCommandHandler) Handle(ctx context.Context, cmd BillingAppleEvent) (err error) {
	return h.billingService.DistributeAppleEvent(ctx, cmd.payload)
}

func NewBillingAppleCommandHandler(
	billingService services.BillingService,
	logger *logrus.Entry,
	metricsClient decorator.MetricsClient,
) decorator.CommandHandler[BillingAppleEvent] {
	return decorator.ApplyCommandDecorators[BillingAppleEvent](
		billingAppleCommandHandler{billingService: billingService},
		logger,
		metricsClient,
	)
}
