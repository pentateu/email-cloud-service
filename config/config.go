package config

import (
	iface "github.com/ipfs/interface-go-ipfs-core"
	"github.com/spf13/cobra"
)

type MailConfig struct {
	mailAccounts []string
}

//Load - load the mail service config
func Load(cmd *cobra.Command, args []string, ipfs iface.CoreAPI) (*MailConfig, error) {

	return nil, nil
}
