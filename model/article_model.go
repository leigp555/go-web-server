package model

import (
	"go/note/util"
	"gorm.io/gorm"
	"log"
)

type Article struct {
	Title   string
	Content string
	gorm.Model
}

func init() {
	var db = util.Mydb.Db
	err := db.AutoMigrate(&User{})
	if err != nil {
		log.Panicln("表创建失败")
	}
	a1 := Article{Title: "七米", Content: "放过"}
	a2 := Article{Title: "沙河娜扎", Content: "只是"}
	db.Create(&a1)
	db.Create(&a2)
}
