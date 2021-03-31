package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"

	"alfred/pkg"
)

var decodeCmd = &cobra.Command{
	Use:   "decode",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			return
		}

		defer func() {
			wf.SendFeedback()
			return
		}()

		base64Decode(strings.Join(args, " "))
	},
}

func base64Decode(input string) {
	secs := fmt.Sprintf("%s", pkg.Base64Decode(input))
	wf.NewItem(secs).
		Subtitle("base64解码").
		Arg(secs).
		Valid(true)
}
