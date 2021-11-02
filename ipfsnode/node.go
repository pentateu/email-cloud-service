package ipfs

import (
	"context"
	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/ipfs/go-ipfs/core"
	"github.com/ipfs/go-ipfs/core/coreapi"
	"github.com/ipfs/go-ipfs/core/node/libp2p"
	"github.com/ipfs/go-ipfs/plugin/loader"
	"github.com/ipfs/go-ipfs/repo/fsrepo"
	iface "github.com/ipfs/interface-go-ipfs-core"
	"github.com/ipfs/interface-go-ipfs-core/options"
)

type CfgOpt func(*config.Config)

//spawn - attempts to open a ipfs node based on the configuration root path.
//if fails try to open a node on a temprary folder.
func spawn(ctx context.Context) (iface.CoreAPI, error) {
	defaultPath, err := config.PathRoot()
	if err != nil {
		// shouldn't be possible
		return nil, err
	}

	if err := setupPlugins(defaultPath); err != nil {
		return nil, err
	}

	ipfs, err := open(ctx, defaultPath)
	if err == nil {
		return ipfs, nil
	}

	return tmpNode(ctx)
}

// setupPlugins - Load ipfs plugins from the folder {path}/plugins/
func setupPlugins(path string) error {
	plugins, err := loader.NewPluginLoader(filepath.Join(path, "plugins"))
	if err != nil {
		return fmt.Errorf("error loading plugins: %s", err)
	}

	if err := plugins.Initialize(); err != nil {
		return fmt.Errorf("error initializing plugins: %s", err)
	}

	if err := plugins.Inject(); err != nil {
		return fmt.Errorf("error initializing plugins: %s", err)
	}

	return nil
}

//open - open a ipfs node for a given folder.
func open(ctx context.Context, repoPath string) (iface.CoreAPI, error) {
	// Open the repo
	r, err := fsrepo.Open(repoPath)
	if err != nil {
		return nil, err
	}

	// Construct the node
	node, err := core.NewNode(ctx, &core.BuildCfg{
		Online:  true,
		Routing: libp2p.DHTClientOption,
		Repo:    r,
	})
	if err != nil {
		return nil, err
	}
	return coreapi.NewCoreAPI(node)
}

//temp load pluigins and creates a temporary node
func temp(ctx context.Context) (iface.CoreAPI, error) {
	defaultPath, err := config.PathRoot()
	if err != nil {
		// shouldn't be possible
		return nil, err
	}

	if err := setupPlugins(defaultPath); err != nil {
		return nil, err
	}

	return tmpNode(ctx)
}

//tmpNode - creates a temporary node 'dhtclient' on a temp folder.
func tmpNode(ctx context.Context) (iface.CoreAPI, error) {
	dir, err := ioutil.TempDir("", "ipfs-shell")
	if err != nil {
		return nil, fmt.Errorf("failed to get temp dir: %s", err)
	}

	// Cleanup temp dir on exit
	//TODO: add clean up function
	// addCleanup(func() error {
	// 	return os.RemoveAll(dir)
	// })

	identity, err := config.CreateIdentity(ioutil.Discard, []options.KeyGenerateOption{
		options.Key.Type(options.Ed25519Key),
	})
	if err != nil {
		return nil, err
	}
	cfg, err := config.InitWithIdentity(identity)
	if err != nil {
		return nil, err
	}

	// configure the temporary node
	cfg.Routing.Type = "dhtclient"

	err = fsrepo.Init(dir, cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to init ephemeral node: %s", err)
	}
	return open(ctx, dir)
}
