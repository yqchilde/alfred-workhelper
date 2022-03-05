package pkg

import (
	"crypto/sha256"
	"encoding/hex"
	"math/rand"

	"golang.org/x/crypto/pbkdf2"
)

// Vector is used for generate password
type Vector struct {
	password string
	salt     string
	iter     int
	length   int
}

// Bytes2Password generate password
func Bytes2Password(bytes []byte) string {
	h := sha256.New()
	h.Write(bytes)
	md := h.Sum(nil)
	return hex.EncodeToString(md)
}

// EncryptPassword encrypt password
func EncryptPassword(password string, salt string) string {
	var v Vector
	v.password = password
	v.salt = salt
	v.iter = 4096
	v.length = 25

	bytes := pbkdf2.Key([]byte(v.password), []byte(v.salt), v.iter, v.length, sha256.New)
	return "PBKDF2" + ";" + salt + ";" + Bytes2Password(bytes)
}

func bytes2String(bytes []byte) string {
	for i := 0; i < len(bytes); i++ {
		bytes[i] = bytes[i]%93 + '!'
		if bytes[i] == ';' {
			bytes[i] = byte(i) + byte(1) + ';'
		}
	}
	return string(bytes)
}

// GenerateSalt return a salt
func GenerateSalt(length int) string {
	bytes := make([]byte, length)
	_, err := rand.Read(bytes)
	if err != nil {
		return ""
	}

	return bytes2String(bytes)
}
