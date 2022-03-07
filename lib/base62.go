package lib

// 将自增数字10进制转化成62进制
import (
	"math"
	"strings"
)

const Code62 = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func Encode(num int64) string {
	bytes := []byte{}
	for num > 0 {
		bytes = append(bytes, Code62[num%62])
		num = num / 62
	}
	reverse(bytes)
	return string(bytes)
}

func Decode(str string) int64 {
	var num int64
	n := len(str)
	for i := 0; i < n; i++ {
		pos := strings.IndexByte(Code62, str[i])
		num += int64(math.Pow(62, float64(n-1-i)) * float64(pos))
	}
	return num
}

func reverse(b []byte) {
	for left, right := 0, len(b)-1; left < right; left, right = left+1, right-1 {
		b[left], b[right] = b[right], b[left]
	}
}
