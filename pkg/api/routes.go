package api

import (
	"net/http"

	"github.com/dstm45/vacances/pkg/services"
)

func NewRoutes() *http.ServeMux {
	userService := services.NewUserService()
	mux := http.NewServeMux()
	return mux
}
