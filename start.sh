#!/bin/bash

cd /usr/src/code/mall/service/user/rpc
go run user.go -f etc/user.yaml >service.log &
cd /usr/src/code/mall/service/user/api
go run user.go -f etc/user.yaml >service.log &


cd /usr/src/code/mall/service/product/rpc
go run product.go -f etc/product.yaml >service.log &
cd /usr/src/code/mall/service/product/api
go run product.go -f etc/product.yaml >service.log &


cd /usr/src/code/mall/service/order/rpc
go run order.go -f etc/order.yaml >service.log &
cd /usr/src/code/mall/service/order/api
go run order.go -f etc/order.yaml >service.log &

cd /usr/src/code/mall/service/pay/rpc
go run pay.go -f etc/pay.yaml >service.log &
cd /usr/src/code/mall/service/pay/api
go run pay.go -f etc/pay.yaml >service.log &