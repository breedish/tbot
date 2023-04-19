package services

import (
	"context"
	"math/rand"
	"testing"
	"time"

	"github.com/breedish/tbot/internal/bot/ports/out/gateway"
	"github.com/stretchr/testify/assert"
)

func TestBillingService(t *testing.T) {
	t.Parallel()
	rand.Seed(time.Now().UTC().UnixNano())

	t.Run("testDistributeBillingAppleEvent", func(t *testing.T) {
		t.Parallel()
		testDistributeBillingAppleEvent(t)
	})
}

type BillingGatewayMock struct {
	success map[string]bool
}

func (m BillingGatewayMock) DispatchAppleEvent(payload []byte, target string) bool {
	var _, present = m.success[target]
	return present
}

func testDistributeBillingAppleEvent(t *testing.T) {
	t.Helper()

	testCases := []struct {
		Name    string
		gateway gateway.BillingGateway
		payload string
		result  bool
	}{
		{
			Name:    "should return success for live mode event when event was processed",
			gateway: BillingGatewayMock{success: map[string]bool{"https://everymeet.com": true}},
			payload: "{mode:Live}",
			result:  true,
		},
		{
			Name:    "should return failure for live mode event when event was not processed by any target",
			gateway: BillingGatewayMock{success: map[string]bool{}},
			payload: "{mode:Live}",
			result:  false,
		},
		{
			Name:    "should return failure for sandbox mode event when event was not processed by any target",
			gateway: BillingGatewayMock{success: map[string]bool{}},
			payload: "{mode:Sandbox}",
			result:  false,
		},
		{
			Name:    "should return true for sandbox mode event when event was processed by single target",
			gateway: BillingGatewayMock{success: map[string]bool{"https://test.everymeet.com": true}},
			payload: "{mode:Sandbox}",
			result:  true,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()

			billingService := NewBillingService(tc.gateway)
			result := billingService.DistributeAppleEvent(context.TODO(), []byte(tc.payload))

			assert.EqualValues(t, tc.result, result == nil, tc.Name)
		})
	}
}
