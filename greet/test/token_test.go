package test

import (
	"greet/helper"
	"testing"
)

func TestToken(t *testing.T){
	token,err:=helper.GenerateToken("xiaowei","123456")
	if err!=nil{
		t.Log(err)
	}
	t.Log(token)
	mc,err:=helper.ParseToken(token)
	t.Log(mc)
}
