package cmd

import (
	"github.com/spf13/cobra"

	"alfred/pkg/codec"
)

var decodeCmd = &cobra.Command{
	Use: "decode",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			return
		}

		defer func() {
			wf.SendFeedback()
			return
		}()

		base64Decode(args)
	},
}

func base64Decode(args []string) {
	secs := codec.Base64Decode(args[0])
	wf.NewItem(secs).
		Subtitle("base64解码").
		Arg(secs).
		Valid(true)
}
