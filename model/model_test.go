package model

import "testing"

func TestConnect(t *testing.T) {
	InitMongo()
	test()
}

func TestGet(t *testing.T) {
	InitMongo()
	GetOriginURL("1wwi2")

}
