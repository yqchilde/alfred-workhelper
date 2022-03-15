package sign

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
)

// EncryptAES 加密 aes_128_cbc
func EncryptAES(src []byte, key []byte) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	src = pkcs5Padding(src, block.BlockSize())
	blockMode := cipher.NewCBCEncrypter(block, key)
	encrypted := make([]byte, len(src))
	blockMode.CryptBlocks(encrypted, src)
	return base64.StdEncoding.EncodeToString(encrypted), nil
}

// 填充数据
func pkcs5Padding(src []byte, blockSize int) []byte {
	padNum := blockSize - len(src)%blockSize
	pad := bytes.Repeat([]byte{byte(padNum)}, padNum)
	return append(src, pad...)
}
