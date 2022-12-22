package model

import (
	"fmt"
	"go/note/util"
	"gorm.io/gorm"
	"log"
)

type User struct {
	Name string
	Age  int
	gorm.Model
}

func (User) Generate() {
	var db = util.Mydb.Db
	err := db.AutoMigrate(&User{})
	fmt.Printf("我是db%v", db)
	if err != nil {
		log.Panicln("User表创建失败")
	}
}
