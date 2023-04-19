package in

import (
	"net/http"

	"github.com/breedish/tbot/internal/bot/app/command"

	"github.com/breedish/tbot/internal/bot/app"
	"github.com/breedish/tbot/internal/common/server/httperr"
)

type HttpServer struct {
	application app.ApplicationUseCases
}

func NewHttpServer(app app.ApplicationUseCases) HttpServer {
	return HttpServer{app}
}

func (h HttpServer) DoHealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func (h HttpServer) HandleAppleBillingEvent(w http.ResponseWriter, r *http.Request) {
	cmd := command.BillingAppleEvent{}
	err := h.application.Commands.BillingAppleCommand.Handle(r.Context(), cmd)

	if err != nil {
		httperr.InternalError("processing issue", err, w, r)
		return
	}

	w.WriteHeader(http.StatusOK)
}
