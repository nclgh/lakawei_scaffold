package rpc

import (
	"github.com/nclgh/lakawei_rpc/client"
)

var (
	clients = make(map[string]*client.RpcClient)
)

func GetRpcClient(serviceName string) *client.RpcClient {
	if cli, ok := clients[serviceName]; ok {
		return cli
	}
	cli, err := client.InitClient(serviceName)
	if err != nil {
		panic()
	}
}
