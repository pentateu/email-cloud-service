package ipfs

import (
	"context"

	ipfshttp "github.com/ipfs/go-ipfs-http-client"
	iface "github.com/ipfs/interface-go-ipfs-core"
)

//http - connect the node using http
func http(ctx context.Context) (iface.CoreAPI, error) {
	httpApi, err := ipfshttp.NewLocalApi()
	if err != nil {
		return nil, err
	}
	err = httpApi.Request("version").Exec(ctx, nil)
	if err != nil {
		return nil, err
	}
	return httpApi, nil
}
