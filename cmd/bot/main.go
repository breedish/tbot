package main

import (
	"github.com/breedish/tbot/internal/bot/context"
	"github.com/breedish/tbot/internal/bot/ports/in"
	"github.com/breedish/tbot/internal/common/logs"
	"github.com/breedish/tbot/internal/common/server"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func main() {
	logs.Init()

	application, ctx, cleanup := context.NewApplication()
	defer cleanup()

	server.RunHTTPServer(ctx.ServiceInfo, ctx.MetricsRegistry, func(router chi.Router) http.Handler {
		return in.HandlerFromMux(in.NewHttpServer(application), router)
	})
}
