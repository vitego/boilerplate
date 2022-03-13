package controllers

import (
	"context"
	"github.com/vitego/router/manager"
	"github.com/vitego/router/response"
	"net/http"
)

/*
	@Route("GET", "/ping")
*/
func (Handler) Ping(ctx context.Context, m *manager.Manager, w http.ResponseWriter) bool {
	return response.Success(w, http.StatusOK, map[string]string{"message": "pong"})
}
