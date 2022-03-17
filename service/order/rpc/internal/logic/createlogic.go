package logic

import (
	"context"
	"mall/service/order/model"
	"mall/service/order/rpc/internal/svc"
	"mall/service/order/rpc/order"
	"mall/service/product/rpc/product"
	"mall/service/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/status"
)

type CreateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateLogic {
	return &CreateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateLogic) Create(in *order.CreareRequest) (*order.CreateResponse, error) {
	// todo: add your logic here and delete this line
	// 查询用户是否存在
	_,err := l.svcCtx.UserRpc.UserInfo(l.ctx,&user.UserInfoRequest{
		Id:in.Uid,
	})
	if(err != nil) {
		return nil,err;
	}
	// 查询产品是否存在
	productRes,err := l.svcCtx.ProductRpc.Detail(l.ctx,&product.DetailRequest{
		Id:in.Pid,
	})
	if(err!= nil) {
		return nil,err;
	}
	if(productRes.Stock<=0) {
		return nil, status.Error(500,"产品库存不足");
	}
	// 创建新订单
	newOrder := model.Order{
		Uid: in.Uid,
		Pid: in.Pid,
		Amount: in.Amount,
		Status: 0,
	}
	// 新订单插入数据库
	res,err:= l.svcCtx.OrderModel.Insert(&newOrder)
	if(err!=nil) {
		return nil,status.Error(500,err.Error())
	}

	newOrder.Id,err  = res.LastInsertId()
	if(err!= nil) {
		return nil, status.Error(500,err.Error())
	}

	// 更新库存
	_,err = l.svcCtx.ProductRpc.Update(l.ctx,&product.UpdateRequest{
		Id: productRes.Id,
		Name: productRes.Name,
		Desc: productRes.Desc,
		Stock: productRes.Stock-1,
		Amount: productRes.Amount,
		Status: productRes.Status,
	})
	if(err!=nil) {
		return nil,err
	}

	return &order.CreateResponse{
		Id: newOrder.Id,
	}, nil
}
