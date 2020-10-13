package handler

import (
	"net/http"

	"bookstore/user/internal/logic/profile"
	"bookstore/user/internal/svc"
	"bookstore/user/internal/types"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func GetProfileHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewGetProfileLogic(r.Context(), ctx)
		resp, err := l.GetProfile(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.WriteJson(w, http.StatusOK, resp)
		}
	}
}
