package config

import (
	"github.com/zeromicro/go-zero/zrpc"
	"github.com/zeromicro/go-zero/core/stores/cache"
)

type Config struct {
	zrpc.RpcServerConf
	Mysql struct{
		DataSource string
	}
	CacheRedis cache.CacheConf
	UserRpc zrpc.RpcClientConf
	ProductRpc zrpc.RpcClientConf
	DTM string
}
