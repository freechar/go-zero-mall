package svc

import (
	"mall/service/pay/model"
	"mall/service/pay/rpc/internal/config"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"mall/service/user/rpc/userclient"
	"mall/service/order/rpc/orderclient"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config
	PayModel model.PayModel
	UserRpc userclient.User
	OrderRpc orderclient.Order
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn:=sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config: c,
		PayModel: model.NewPayModel(conn,c.CacheRedis),
		UserRpc:  userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
		OrderRpc: orderclient.NewOrder(zrpc.MustNewClient(c.OrderRpc)),
	}
}
