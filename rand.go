package stringhelper

import (
	"math/rand"
	"time"
)

const (
	RandomStringModNumberPlusLetter           = 1
	RandomStringModNumberPlusLetterPlusSymbol = 2
	RandomStringModNumber                     = 3
)

func GenRandomStr(length int64, mod uint32) string {
	var strKey string
	if mod == RandomStringModNumberPlusLetter {
		strKey = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	} else if mod == RandomStringModNumberPlusLetterPlusSymbol {
		strKey = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ~!@#$%^&*(){}|\\[]?/"
	} else if mod == RandomStringModNumber {
		strKey = "0123456789"
	} else {
		strKey = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	bytes := make([]byte, length)
	for i := range bytes {
		bytes[i] = strKey[r.Intn(len(strKey))]
	}

	return string(bytes)
}
