package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"

	"alfred/pkg"
)

var encryptCmd = &cobra.Command{
	Use:   "encrypt",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			return
		}

		defer func() {
			wf.SendFeedback()
			return
		}()

		encryptString(strings.Join(args, " "))
	},
}

func encryptString(input string) {
	secs := fmt.Sprintf("%s", encryptPassword(input))
	wf.NewItem(secs).
		Subtitle("encrypt编码").
		Arg(secs).
		Valid(true)
}

func encryptPassword(password string) string {
	salt := pkg.GenerateSalt(6)
	encPwd := pkg.EncryptPassword(password, salt)
	return encPwd
}
