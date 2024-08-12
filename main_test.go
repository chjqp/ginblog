package main

import (
	_ "errors"
	"ginblog/model"
	"testing"
)

//var newError error =errors.New("this is a  new error")

func Test(t *testing.T) {
	model.InitDb()
	a := 1
	if a == 1 {
		t.Error("ok")
	} else {
		t.Error("error")
	}
}
