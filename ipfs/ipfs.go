package ipfs

import (
	"context"
	"fmt"

	iface "github.com/ipfs/interface-go-ipfs-core"
	"github.com/spf13/cobra"
)

//Start the IPFS node service
func Start(cmd *cobra.Command, args []string) (iface.CoreAPI, error) {
	//TODO: Grab nodeType from config
	nodeType := "spawn"
	//TODO: Grab peers 	from config
	peers := []string{"", ""}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var ipfs iface.CoreAPI
	var err error
	switch nodeType {
	case "spawn":
		ipfs, err = spawn(ctx)
	case "local":
		ipfs, err = http(ctx)
	case "temp":
		ipfs, err = temp(ctx)
	default:
		return nil, fmt.Errorf("no such 'node' strategy, %q", nodeType)
	}
	if err != nil {
		return nil, err
	}

	go connect(ctx, ipfs, peers)

	return ipfs, nil
}
