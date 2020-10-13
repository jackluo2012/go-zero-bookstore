package logic

import (
	"context"

	"bookstore/user/internal/svc"
	"bookstore/user/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetUserLogic {
	return GetUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserLogic) GetUser(req types.GetRequest) (*types.GetResponse, error) {
	// todo: add your logic here and delete this line

	return &types.GetResponse{}, nil
}
