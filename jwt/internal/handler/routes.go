// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"bookstore/jwt/internal/svc"

	"github.com/tal-tech/go-zero/rest"
)

func RegisterHandlers(engine *rest.Server, serverCtx *svc.ServiceContext) {
	engine.AddRoutes([]rest.Route{
		{
			Method:  http.MethodPost,
			Path:    "/user/token",
			Handler: jwtHandler(serverCtx),
		},
	})

	engine.AddRoutes([]rest.Route{
		{
			Method:  http.MethodPost,
			Path:    "/user/info",
			Handler: getUserHandler(serverCtx),
		},
	}, rest.WithJwt(serverCtx.Config.JwtAuth.AccessSecret))
}
