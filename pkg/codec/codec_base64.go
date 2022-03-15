package codec

import (
	"encoding/base64"
	"fmt"
)

func Base64Encode(data string) string {
	return base64.StdEncoding.EncodeToString([]byte(data))
}

func Base64Decode(data string) string {
	sDec, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		fmt.Printf("Error decoding string: %s ", err.Error())
		return ""
	}

	return string(sDec)
}
