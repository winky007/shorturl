package models

import (
	"crypto/md5"
	"fmt"
	"io"
	"math/rand"
)

func init() {

}

func GetRandStr(i int) string {
	str := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
	return str[i]
}

func GetRand() int {
	r := rand.Intn(36)
	return r
}

func GetMD5(longUrl string) string {
	h := md5.New()
	salt := "salt4shorturl"
	io.WriteString(h, longUrl+salt)
	urlmd5 := fmt.Sprintf("%x", h.Sum(nil))
	return urlmd5
}
