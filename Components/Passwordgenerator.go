package main

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
)

func encryptPassword(password string) string {
	h := sha256.New()
	h.Write([]byte(password))
	b := h.Sum(nil)
	return "{SHA256}" + base64.StdEncoding.EncodeToString(b)
}

func encryptPassword2(password string) string {
	h := sha256.Sum256([]byte(password))
	return "{SHA256}" + base64.StdEncoding.EncodeToString(h[:])
}

func main() {
	fmt.Println(encryptPassword("iniPassword"))
	fmt.Println(encryptPassword2("abcd1234"))
}
