package ipfs

import (
	"context"
	"fmt"
	"net/url"
	gopath "path"

	iface "github.com/ipfs/interface-go-ipfs-core"
	ipath "github.com/ipfs/interface-go-ipfs-core/path"
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

//parsePath - parse an IPFS path and return a Path obj instance.
func parsePath(path string) (ipath.Path, error) {
	ipfsPath := ipath.New(path)
	if ipfsPath.IsValid() == nil {
		return ipfsPath, nil
	}

	u, err := url.Parse(path)
	if err != nil {
		return nil, fmt.Errorf("%q could not be parsed: %s", path, err)
	}

	switch proto := u.Scheme; proto {
	case "ipfs", "ipld", "ipns":
		ipfsPath = ipath.New(gopath.Join("/", proto, u.Host, u.Path))
	case "http", "https":
		ipfsPath = ipath.New(u.Path)
	default:
		return nil, fmt.Errorf("%q is not recognized as an IPFS path", path)
	}
	return ipfsPath, ipfsPath.IsValid()
}
