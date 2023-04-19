package gateway

type BillingGateway interface {
	DispatchAppleEvent(payload []byte, target string) bool
}
