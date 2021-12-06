package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var (
	// pusher service config option
	t2Type string
	t2File string
)
var test2Cmd = &cobra.Command{
	Use:   "t2",
	Short: "权限中心服务",
	Long:  `权限中心服务`,
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("t2")
		fmt.Println(t2Type)
		return nil
	},
}

func init() {
	test2Cmd.Flags().StringVarP(&t2Type, "t2-type", "t", "file", "the service config type [file/env/etcd]")
	test2Cmd.Flags().StringVarP(&t2File, "t2-file", "f", "etc/keyauth.toml", "the service config from file")
	RootCmd.AddCommand(test2Cmd)
}
