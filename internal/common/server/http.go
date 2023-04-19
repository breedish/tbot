package server

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/breedish/tbot/internal/common/logs"

	"github.com/breedish/tbot/internal/common/service"
	"github.com/prometheus/client_golang/prometheus"
	prom_metrics "github.com/slok/go-http-metrics/metrics/prometheus"
	http_metrics "github.com/slok/go-http-metrics/middleware"
	"github.com/slok/go-http-metrics/middleware/std"

	"github.com/breedish/tbot/internal/common/auth"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/httplog"
	"github.com/go-chi/render"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sirupsen/logrus"
)

func RunHTTPServer(serviceInfo service.Info, registry prometheus.Registerer, createHandler func(router chi.Router) http.Handler) {
	apiRouter := chi.NewRouter()
	setMiddlewares(apiRouter, serviceInfo, registry)

	rootRouter := chi.NewRouter()
	rootRouter.Mount(serviceInfo.GetServiceURL(), createHandler(apiRouter))
	rootRouter.Mount(serviceInfo.GetServiceURL()+"/q/metrics", promhttp.Handler())

	logrus.Info(fmt.Sprintf("Server started %s:%s", serviceInfo.GetServiceURL(), serviceInfo.Port))

	instance := &http.Server{
		Addr:              ":" + serviceInfo.Port,
		Handler:           rootRouter,
		ReadHeaderTimeout: 5 * time.Second,
	}

	err := instance.ListenAndServe()
	if err != nil {
		logrus.WithError(err).Panic("Unable to start HTTP server")
	}
}

func setMiddlewares(router *chi.Mux, serviceInfo service.Info, registry prometheus.Registerer) {
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.NoCache)
	router.Use(middleware.Recoverer)
	router.Use(middleware.Timeout(30 * time.Second))
	router.Use(middleware.Heartbeat(serviceInfo.GetServiceURL() + "/ping"))
	router.Use(render.SetContentType(render.ContentTypeJSON))
	router.Use(logs.NewStructuredLogger(logrus.StandardLogger()))
	addMetricsMiddleware(router, serviceInfo, registry)
	addAuthMiddleware(router)
	addAccessRequestMiddleware(router, serviceInfo)
}

func addMetricsMiddleware(router *chi.Mux, serviceInfo service.Info, registry prometheus.Registerer) {
	recorder := prom_metrics.NewRecorder(prom_metrics.Config{Registry: registry})
	router.Use(std.HandlerProvider("", http_metrics.New(http_metrics.Config{
		Recorder: recorder,
		Service:  serviceInfo.ServiceName,
	})))
}

func addAuthMiddleware(router *chi.Mux) {
	if skipAuth, _ := strconv.ParseBool(os.Getenv("SKIP_AUTH")); skipAuth {
		return
	}
	if mockAuth, _ := strconv.ParseBool(os.Getenv("MOCK_AUTH")); mockAuth {
		router.Use(auth.HttpMockMiddleware)
		return
	}
	router.Use(auth.EverymeetAuthMiddleware{}.Middleware)
}

func addAccessRequestMiddleware(router *chi.Mux, serviceInfo service.Info) {
	router.Use(httplog.RequestLogger(httplog.NewLogger(serviceInfo.ServiceName, httplog.Options{JSON: true})))
}
