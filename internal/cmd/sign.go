package cmd

import (
	"encoding/json"
	"strconv"
	"time"

	"github.com/spf13/cobra"

	"alfred/pkg/sign"
)

var signCmd = &cobra.Command{
	Use: "sign",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			return
		}

		defer func() {
			wf.SendFeedback()
			return
		}()

		encryptionSignUsePBKDF2(args)
		encryptionSignUseAES(args)

	},
}

// pbkdf2算法加密
func encryptionSignUsePBKDF2(args []string) {
	secs := sign.EncryptPBKDF2(args[0])
	wf.NewItem(secs).
		Subtitle("Sign --> PBKDF2").
		Arg(secs).
		Valid(true)
}

// aes-128-cbc + base64加密
func encryptionSignUseAES(args []string) {
	if len(args) != 2 {
		return
	}

	type register struct {
		PaasID int
		AppID  int
		Time   int64
	}
	paasID, err := strconv.Atoi(args[0])
	if err != nil {
		return
	}

	data, _ := json.Marshal(register{
		PaasID: paasID,
		AppID:  1,
		Time:   time.Now().Unix(),
	})
	signStr, err := sign.EncryptAES(data, []byte(args[1]))
	if err != nil {
		return
	}

	wf.NewItem(signStr).
		Subtitle("Sign --> AES-128-CBC").
		Arg(signStr).
		Valid(true)
}
