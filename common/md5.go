package common

import (
	"crypto/md5"
	"encoding/hex"
)

func GetMD5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}

func CheckPasswrd(check, password string) bool {
	return GetMD5Hash(check) == password
}