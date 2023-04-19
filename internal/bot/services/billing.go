package services

import (
	"context"
	"strings"
	"sync"

	"github.com/breedish/tbot/internal/bot/ports/out/gateway"

	"github.com/breedish/tbot/internal/common/errors"
)

type BillingService interface {
	DistributeAppleEvent(ctx context.Context, payload []byte) (err error)
}

type billingService struct {
	gateway gateway.BillingGateway
}

var liveModeTargets = []string{"https://everymeet.com"}
var sandboxModeTargets = []string{"https://everymeet.com", "https://dev1.everymeet.com", "https://test.everymeet.com"}

func (s billingService) DistributeAppleEvent(ctx context.Context, payload []byte) (err error) {
	body := string(payload)

	var targets []string
	if strings.Contains(body, "Sandbox") {
		targets = sandboxModeTargets
	} else {
		targets = liveModeTargets
	}

	wg := sync.WaitGroup{}
	wg.Add(len(targets))

	responses := make(chan bool, len(targets))

	for _, target := range targets {
		go func(t string) {
			defer wg.Done()
			responses <- s.gateway.DispatchAppleEvent(payload, t)
		}(target)
	}

	wg.Wait()
	close(responses)

	for res := range responses {
		if res {
			return nil
		}
	}

	return errors.NewProcessingError("issue processing apple event", "billing-event-handler-issue")
}

func NewBillingService(gateway gateway.BillingGateway) BillingService {
	return billingService{gateway: gateway}
}
