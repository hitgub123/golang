package dao

import (
	// "fmt"
	// "gorm.io/driver/postgres"
	// "gorm.io/gorm"
	"log"
	"slq.me/obj"
)

func PageUser(p obj.Pagination, keyword string) *obj.Result {
	db := GetDB()
	users := []obj.User{}
	if keyword != "" {
		db = db.Where("name like ?", "%"+keyword+"%")
	}
	err := db.Limit(p.PageSize).Offset((p.PageNum - 1) * p.PageSize).Order("id desc").Find(&users).Error
	if err == nil {
		return &obj.Result{Code: 0, Error: nil, Data: &users}
	} else {
		return &obj.Result{Code: 1, Error: err, Data: nil}
	}
}

func OneUser2(username string, password string) *obj.Result {
	db := GetDB()
	var user *obj.User
	err := db.Where("name = ? and password=?", username, password).First(&user).Error
	log.Println("OneUser2 dao>>",user)
	if err == nil {
		return &obj.Result{Code: 0, Error: nil, Data: &user}
	} else {
		return &obj.Result{Code: 1, Error: err, Data: nil}
	}
}

func OneUser(id string) *obj.Result {
	db := GetDB()
	var user *obj.User
	err := db.Where("id = ?", id).First(&user).Error
	// db.Unscoped().First(&user, id)
	// db.First(&user, id)
	log.Println("OneUser>>", err)
	if err == nil {
		return &obj.Result{Code: 0, Error: nil, Data: &user}
	} else {
		return &obj.Result{Code: 1, Error: err, Data: nil}
	}
}

func DeleteUser(id string) *obj.Result {
	db := GetDB()
	var user *obj.User
	// db.Unscoped().Where("id = ?", id).Delete(&user)
	err := db.Where("id = ?", id).Delete(&user).Error
	log.Println("DeleteUser>>", err)
	if err == nil {
		return &obj.Result{Code: 0, Error: nil, Data: nil}
	} else {
		return &obj.Result{Code: 1, Error: err, Data: nil}
	}
}

func UpdateUser(o obj.User) *obj.Result {
	// log.Println("dao1",o)
	db := GetDB()
	id := o.Id
	o.Id = 0
	// log.Println("dao2",o)
	err := db.Model(&obj.User{}).Where("id = ?", id).Updates(&o).Error
	log.Println("UpdateUser>>", err)
	if err == nil {
		return &obj.Result{Code: 0, Error: nil, Data: nil}
	} else {
		return &obj.Result{Code: 1, Error: err, Data: nil}
	}
}

func InsertUser(o obj.User) *obj.Result {
	o.Id = 0
	err := GetDB().Create(&o).Error
	log.Println("InsertUser>>", err,o)
	if err == nil {
		return &obj.Result{Code: 0, Error: nil, Data: nil}
	} else {
		return &obj.Result{Code: 1, Error: err, Data: nil}
	}
}
