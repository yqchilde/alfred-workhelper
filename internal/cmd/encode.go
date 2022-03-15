package cmd

import (
	"github.com/spf13/cobra"

	"alfred/pkg/codec"
)

var encodeCmd = &cobra.Command{
	Use: "encode",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			return
		}

		defer func() {
			wf.SendFeedback()
			return
		}()
		base64Encode(args)
		md5Encode(args)
	},
}

func base64Encode(args []string) {
	secs := codec.Base64Encode(args[0])
	wf.NewItem(secs).
		Subtitle("base64编码").
		Arg(secs).
		Valid(true)
}

func md5Encode(args []string) {
	secs := codec.MD5Encode(args[0])
	wf.NewItem(secs).
		Subtitle("md5编码").
		Arg(secs).
		Valid(true)
}
