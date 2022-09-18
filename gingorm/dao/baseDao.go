package dao

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"slq.me/obj"
)

var db *gorm.DB
var dsn = "host=localhost user=postgres password=kimoji dbname=study port=5432 sslmode=disable TimeZone=Asia/Shanghai"

func GetDB() *gorm.DB {
	var err error
	if db == nil {
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			panic(err)
		}
	}
	return db
}

func Test1() {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	if err = db.AutoMigrate(&obj.User{}); err != nil {
		panic(err)
	}
	fmt.Println("migrated")
	a := &obj.User{Name: "测试",}
	result := db.Create(a)
	log.Println(a.Id)                // 返回插入数据的主键
	log.Println(result.Error)        // 返回 error
	log.Println(result.RowsAffected) // 返回插入记录的条数
	//批量插入
	b := []obj.User{{Name: "测试1", }, {Name: "测试2",},}
	result1 := db.Create(&b)
	log.Println(b[0].Id)              // 返回插入数据的主键
	log.Println(result1.Error)        // 返回 error
	log.Println(result1.RowsAffected) // 返回插入记录的条数

}
