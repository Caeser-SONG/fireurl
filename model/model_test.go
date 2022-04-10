package model

import (
	"fmt"
	"testing"
)

func rrTestConnect(t *testing.T) {
	InitMongo()
	test()
}

func rrTestGet(t *testing.T) {
	InitMongo()
	GetOriginURL("1wwi2")

}

func TestExist(t *testing.T) {
	InitMongo()
	fmt.Println(IsExist("aaswws"))
	fmt.Println(IsExist("6f698"))

}
