package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"alfred/pkg/net"
)

var netCmd = &cobra.Command{
	Use: "ip",
	Run: func(cmd *cobra.Command, args []string) {
		defer func() {
			wf.SendFeedback()
			return
		}()

		getInterIP()
	},
}

func getInterIP() {
	ips := net.GetInterIP()
	for i, secs := range ips {
		wf.NewItem(secs).
			Subtitle(fmt.Sprintf("IP地址%d: %s", i+1, secs)).
			Arg(secs).
			Valid(true)
	}
}
