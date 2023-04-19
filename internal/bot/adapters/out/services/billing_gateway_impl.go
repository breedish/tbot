package services

import (
	"bytes"
	"context"
	"net/http"
)

type BillingGatewayImpl struct {
	Client http.Client
}

func (gateway BillingGatewayImpl) DispatchAppleEvent(payload []byte, target string) bool {
	req, err := http.NewRequestWithContext(context.TODO(), "POST", target+"/api/billing/events/aapl", bytes.NewBuffer(payload))
	if err != nil {
		return false
	}
	req.Header.Set("Content-Type", "application/json")

	res, err := gateway.Client.Do(req)
	if err != nil {
		return false
	}

	defer res.Body.Close()

	return res.StatusCode == http.StatusOK
}
