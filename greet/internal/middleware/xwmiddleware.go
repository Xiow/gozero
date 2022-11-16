package middleware

import (
	"context"
	"fmt"
	"greet/helper"
	"greet/model"
	"net/http"
)

type XwMiddleware struct {
}

func NewXwMiddleware() *XwMiddleware {
	return &XwMiddleware{}
}

func (m *XwMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO generate middleware implement function, delete after code implementation

		name:=r.FormValue("name")
		password:=r.FormValue("password")
		token,err:=helper.GenerateToken(name,password)
		if err!=nil{
			fmt.Println("放入信息到redis失败")
			return
		}
		rdb:=model.InitRedisClient()
		_,err=rdb.Do(context.Background(),"set",name,token).Result()
		if err!=nil{
			fmt.Println("放入信息到redis失败")
			return
		}
		//password:=http.Request.FormValue(r,"password")
		fmt.Println(name,password)
		// Passthrough to next handler if need
		next(w, r)
	}
}
