package main

import (
	"github.com/pentateu/email-cloud-service/config"
	"github.com/pentateu/email-cloud-service/ipfs"
	"github.com/pentateu/email-cloud-service/mail"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "mail",
	Short: "SMTP server using maildir, ipfs and encryption",
	Long:  `It's a small SMTP server (go-guerrilla) to save encrypted email on ipfs using the maildir standard`,
	Run: func(cmd *cobra.Command, args []string) {
		ipfsNode, err := ipfs.Start(cmd, args)
		if err != nil {
			//TODO: log fatal error :( - ipfs could not start
			return
		}

		mailConfig, err := config.Load(cmd, args, ipfsNode)
		if err != nil {
			//TODO: log fatal error :( - config could not load
			return
		}

		mail.Start(cmd, args, mailConfig, ipfsNode)
	},
}

var (
	verbose bool
)

func init() {
	cobra.OnInitialize()
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false,
		"print out more debug information")
	rootCmd.PersistentPreRun = func(cmd *cobra.Command, args []string) {
		if verbose {
			logrus.SetLevel(logrus.DebugLevel)
		} else {
			logrus.SetLevel(logrus.InfoLevel)
		}
	}
	mail.Init(rootCmd)
}
