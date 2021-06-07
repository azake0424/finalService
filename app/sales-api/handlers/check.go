package handlers

import (
	"context"
	"errors"
	"finalService/foundation/web"
	"log"
	"math/rand"
	"net/http"
)

type check struct {
	log *log.Logger
}

func (c check) readiness(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	if n := rand.Intn(100); n%100 == 0 {
		return web.NewRequestError(errors.New("trusrted error"), http.StatusBadRequest)
	}

	status := struct {
		Status string
	}{
		Status: "Ok",
	}
	return web.Respond(ctx, w, status, http.StatusOK)

}
