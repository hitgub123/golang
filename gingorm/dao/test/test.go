package main

import (
	"slq.me/dao"
	"log"
	"reflect"
	"slq.me/obj"
)

func main() {
	// db:=dao.GetDB()
	u := obj.User{}
	log.Println("type:", reflect.TypeOf(u))
	dao.Test1()

	return
}
