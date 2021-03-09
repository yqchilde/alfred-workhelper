package pkg

import (
	"crypto/md5"
	"fmt"
)

func MD5Encode(data string) string {
	b := []byte(data)
	has := md5.Sum(b)
	return fmt.Sprintf("%x", has)
}
