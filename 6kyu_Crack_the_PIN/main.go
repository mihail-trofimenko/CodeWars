package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

func GetMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

func Crack(hash string) string {

	for pin := 0; pin <= 99999; pin++ {
		if hash == GetMD5Hash(fmt.Sprintf("%05d", pin)) {
			return string(fmt.Sprintf("%05d", pin))
		}
	}
	return ""
}

func main() {
	pin := "12345"
	fmt.Printf("PIN:%s \nHash:%s\n", pin, GetMD5Hash(pin))
	fmt.Println("Cracked PIN:", Crack(GetMD5Hash(pin)))
}
