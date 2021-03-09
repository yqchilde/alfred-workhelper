package service

import (
	"fmt"
	"strings"

	"alfred/pkg"
)

func RunDecode() {
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

	base64Decode(input)

}

func base64Decode(input string) {
	secs := fmt.Sprintf("%s", pkg.Base64Decode(input))
	wf.NewItem(secs).
		Subtitle("base64解码").
		Arg(secs).
		Icon(iconUser).
		Valid(true)
}
