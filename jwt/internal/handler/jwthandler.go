package handler

import (
	"net/http"

	"bookstore/jwt/internal/logic"
	"bookstore/jwt/internal/svc"
	"bookstore/jwt/internal/types"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func jwtHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.JwtTokenRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewJwtLogic(r.Context(), ctx)
		resp, err := l.Jwt(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.WriteJson(w, http.StatusOK, resp)
		}
	}
}
