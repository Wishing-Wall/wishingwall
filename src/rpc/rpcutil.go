package rpcutil

import (
	"github.com/Wishing-Wall/btcrpcclient"
)

var gClient *Client

func init() {
	gClient, err := btcrpcclient.New(&btcrpcclient.ConnConfig{
		HttpPostMode: true,
		DisableTLS:   true,
		Host:         "127.0.0.1:18001",
		User:         "dogecoin",
		Pass:         "DL2rnNd5KnA5HPFW9J84upxyXXRN9yoe3D",
	}, nil)
	if err != nil {
		logger.Errorln("error creating new btc client: %v", err)
	}
	return
}

func GetBlockCount()uint64{
	return gClient.GetBlockCount()
}