package lib

import (
	"fmt"
	"testing"
)

func TestTrans(t *testing.T) {
	var tt *TransMD5
	a := tt.Trans2Short("asdasdasd")
	fmt.Println(a)
}

func TestBase62(t *testing.T) {
	var tt2 *TransBase62
	fmt.Println(tt2.Trans2Short("ssssss"))
}
