package model

import (
	"go/note/util"
	"gorm.io/gorm"
	"log"
)

type User struct {
	Name string
	Age  int
	gorm.Model
}

func init() {
	var db = util.Mydb.Db
	err := db.AutoMigrate(&User{})
	if err != nil {
		log.Panicln("表创建失败")
	}
	u1 := User{Name: "七米", Age: 18}
	u2 := User{Name: "沙河娜扎", Age: 22}
	db.Create(&u1)
	db.Create(&u2)
}
