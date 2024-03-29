package logic

import (
	"context"

	"mall/service/order/api/internal/svc"
	"mall/service/order/api/internal/types"
	"mall/service/order/rpc/orderclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) CreateLogic {
	return CreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateLogic) Create(req types.CreateRequest) (resp *types.CreateResponse, err error) {
	// todo: add your logic here and delete this line
	res,err:= l.svcCtx.OrderRpc.Create(l.ctx,&orderclient.CreareRequest{
		Uid:req.Uid,
		Pid: req.Pid,
		Amount: req.Amount,
		Status: req.Status,
	})
	if err != nil {
		return nil,err
	}
	return &types.CreateResponse{
		Id: res.Id,
	},nil
}
