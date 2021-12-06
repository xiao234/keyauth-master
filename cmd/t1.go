package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var test1Cmd = &cobra.Command{
	Use:   "t1",
	Short: " ",
	Long:  `权限中心服务`,
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("t1")
		return nil
	},
}

func init() {
	//test1Cmd.Flags().StringVarP(&confType, "config-type", "t", "file", "the service config type [file/env/etcd]")
	//test1Cmd.Flags().StringVarP(&confFile, "config-file", "f", "etc/keyauth.toml", "the service config from file")
	RootCmd.AddCommand(test1Cmd)
}
