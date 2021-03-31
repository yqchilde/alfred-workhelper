package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"

	"alfred/pkg"
)

var encodeCmd = &cobra.Command{
	Use:   "encode",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			return
		}

		defer func() {
			wf.SendFeedback()
			return
		}()
		base64Encode(strings.Join(args, " "))
		md5Encode(strings.Join(args, " "))
	},
}

func base64Encode(input string) {
	secs := fmt.Sprintf("%s", pkg.Base64Encode(input))
	wf.NewItem(secs).
		Subtitle("base64编码").
		Arg(secs).
		Valid(true)
}

func md5Encode(input string) {
	secs := fmt.Sprintf("%s", pkg.MD5Encode(input))
	wf.NewItem(secs).
		Subtitle("md5编码").
		Arg(secs).
		Valid(true)
}
