package service

import (
	"slq.me/dao"
	"slq.me/obj"
	// "log"
)

func PageUser(p obj.Pagination,keyword string) *obj.Result{
	return dao.PageUser(p,keyword)

}

func OneUser(id string) *obj.Result{
	return dao.OneUser(id)
}

func OneUser2(username string,password string) *obj.Result{
	return dao.OneUser2(username,password)
}

func DeleteUser(id string) *obj.Result{
	return dao.DeleteUser(id)
}

func UpdateUser(o obj.User) *obj.Result{
	return dao.UpdateUser(o)
}

func InsertUser(o obj.User) *obj.Result{
	return dao.InsertUser(o)
}