package logic

import (
	"context"

	"bookstore/user/internal/svc"
	"bookstore/user/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetProfileLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetProfileLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetProfileLogic {
	return GetProfileLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetProfileLogic) GetProfile(req types.GetRequest) (*types.GetResponse, error) {
	// todo: add your logic here and delete this line

	return &types.GetResponse{}, nil
}
