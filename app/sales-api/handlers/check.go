package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
)

type check struct {
	logger *log.Logger
}

func (c check) readiness(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	r.Context()
	status := struct {
		Status string
	}{
		Status: "Ok",
	}
	c.logger.Println(status)
	return json.NewEncoder(w).Encode(status)

}
