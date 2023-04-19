package context

import (
	"net/http"
	"time"

	"github.com/breedish/tbot/internal/bot/app"

	adaptersservices "github.com/breedish/tbot/internal/bot/adapters/out/services"
	"github.com/breedish/tbot/internal/common/metrics"

	"github.com/breedish/tbot/internal/bot/app/command"
	"github.com/breedish/tbot/internal/bot/services"
	"github.com/breedish/tbot/internal/common/decorator"
	"github.com/breedish/tbot/internal/common/service"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/sirupsen/logrus"
)

type ApplicationContext struct {
	ServiceInfo     service.Info
	MetricsRegistry prometheus.Registerer
	MetricsClient   decorator.MetricsClient
	Logger          *logrus.Entry
	HttpClient      http.Client
	ServicesContext ServicesContext
}

type ServicesContext struct {
	BillingService services.BillingService
}

func NewApplication() (app.ApplicationUseCases, ApplicationContext, func()) {
	var logger = logrus.NewEntry(logrus.StandardLogger())
	var metricsRegistry = prometheus.DefaultRegisterer
	var httpClient = http.Client{Timeout: 30 * time.Second}

	var serviceInfo = service.Info{
		ServiceName: "bot",
		Port:        "8080",
		ApiVersion:  "v2",
	}

	ctx := ApplicationContext{
		ServiceInfo:     serviceInfo,
		MetricsRegistry: metricsRegistry,
		MetricsClient:   metrics.NewApplicationMetricsCollector(serviceInfo, metricsRegistry),
		HttpClient:      httpClient,
		Logger:          logger,
		ServicesContext: ServicesContext{
			BillingService: services.NewBillingService(adaptersservices.BillingGatewayImpl{Client: httpClient}),
		},
	}

	return newApplication(ctx), ctx, func() {}
}

func newApplication(ctx ApplicationContext) app.ApplicationUseCases {
	return app.ApplicationUseCases{
		Commands: app.Commands{
			BillingAppleCommand: command.NewBillingAppleCommandHandler(ctx.ServicesContext.BillingService, ctx.Logger, ctx.MetricsClient),
		},
		Queries: app.Queries{},
	}
}
