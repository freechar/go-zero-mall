package svc

import (
	"mall/service/product/rpc/internal/config"
	"mall/service/product/model"
	"github.com/zeromicro/go-zero/core/stores/sqlx"

)


type ServiceContext struct {
	Config config.Config
	ProductModel model.ProductModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	// 链接musql
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config: c,
		ProductModel: model.NewProductModel(conn,c.CacheRedis),
	}
}
