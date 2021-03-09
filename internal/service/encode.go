package service

import (
	"fmt"
	"strings"

	aw "github.com/deanishe/awgo"

	"alfred/pkg"
)

var (
	iconUser = &aw.Icon{
		Value: aw.IconUser.Value,
		Type:  aw.IconUser.Type,
	}
)

func RunEncode() {
	var err error

	args := wf.Args()

	if len(args) == 0 {
		return
	}

	defer func() {
		if err == nil {
			wf.SendFeedback()
			return
		}
	}()

	input := strings.Join(args, " ")

	base64Encode(input)
	md5Encode(input)

}

func base64Encode(input string) {
	secs := fmt.Sprintf("%s", pkg.Base64Encode(input))
	wf.NewItem(secs).
		Subtitle("base64编码").
		Arg(secs).
		Icon(iconUser).
		Valid(true)
}

func md5Encode(input string) {
	secs := fmt.Sprintf("%s", pkg.MD5Encode(input))
	wf.NewItem(secs).
		Subtitle("md5编码").
		Arg(secs).
		Icon(iconUser).
		Valid(true)
}
