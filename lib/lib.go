package lib

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"math/rand"
	"strconv"
)

type Transer interface {
	Trans2Short(string) string
}

type TransBase62 struct{}

type TransMD5 struct{}

// md5 法
//
//
func (t *TransMD5) Trans2Short(originURL string) (shortURL string) {

	salt := strconv.Itoa(rand.Intn(100))
	// 加一个随机数
	newurl := originURL + salt
	h := md5.New()
	s := h.Sum([]byte(newurl))
	s2 := hex.EncodeToString(s)
	var buffer bytes.Buffer
	buffer.WriteString(string(s2[0]))
	buffer.WriteString(string(s2[11]))
	buffer.WriteString(string(s2[22]))
	buffer.WriteString(string(s2[33]))
	buffer.WriteString(string(s2[45]))
	shortURL = buffer.String()
	return
}

// todo   需要获取自增序号 mysql  redis  or other
func (t *TransBase62) Trans2Short(originURL string) (shortURL string) {
	var num int64 = 2
	return Encode(num)
}
