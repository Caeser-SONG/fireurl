package lib

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"math/rand"
	"strconv"
)

type Transer interface {
	Trans2Short()
}

type TransBase62 struct{}

type TransMD5 struct{}

// md5 法
//
//
func (t *TransMD5) Trans2Short(originURL string) (shortURL string) {
	salt := strconv.Itoa(rand.Intn(100))
	// 加一个随机数
	or := originURL + salt
	h := md5.New()
	s := h.Sum([]byte(originURL))
	s2 := hex.EncodeToString(s)
	fmt.Println(or)
	fmt.Println(s2)
	return "aaaa"
	//todo
}

func (t *TransBase62) Trans2Short(originURL string) (shortURL string) {
	//todo
	// get id
	url := Encode(111111)
	return "zz"
}
