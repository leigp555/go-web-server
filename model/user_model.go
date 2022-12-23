package model

import (
	"fmt"
	"go/note/util"
	"gorm.io/gorm"
	"log"
)

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(20);comment:用户名"`
	Password string `gorm:"type:varchar(50);comment:密码"`
	Email    string `gorm:"type:varchar(20);unique_index;comment:邮箱"`
}

func (User) Generate() {
	var db = util.Mydb.Db
	err := db.AutoMigrate(&User{})
	fmt.Printf("我是db%v", db)
	if err != nil {
		log.Panicln("User表创建失败")
	}
}
