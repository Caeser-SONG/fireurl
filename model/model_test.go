package model

import "testing"

func TestConnect(t *testing.T) {
	InitMongo()
	test()
}

func TestGet(t *testing.T) {
	InitMongo()
	GetURL("1wwi2")
}
