package logic

import (
	"context"

	"bookstore/rpc/add/internal/svc"
	add "bookstore/rpc/add/pb"
	"bookstore/rpc/model"

	"github.com/tal-tech/go-zero/core/logx"
)

type AddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddLogic {
	return &AddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddLogic) Add(in *add.AddReq) (*add.AddResp, error) {
	// todo: add your logic here and delete this line

	//往数据库里面添加代码
	_, err := l.svcCtx.Model.Insert(model.Book{
		Book:  in.Book,
		Price: in.Price,
	})
	//如果 没有添加成功，就报错
	if err != nil {
		return nil, err
	}
	return &add.AddResp{
		Ok: true,
	}, nil
}
