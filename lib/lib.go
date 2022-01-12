package lib

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"math/rand"
)

// md5 法
func Trans2short(originURL string) (shortURL string) {
	salt := string(rand.Intn(100))
	// 加一个随机数
	or := originURL + salt
	h := md5.New()
	s := h.Sum([]byte(originURL))
	s2 := hex.EncodeToString(s)

	fmt.Println(s2)
	return "aaaa"
	//todo
}
