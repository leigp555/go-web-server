package model

import (
	"go/note/util"
	"gorm.io/gorm"
	"log"
)

type Article struct {
	gorm.Model
	Title   string
	Content string
}

// Generate 创建文章表
func (Article) Generate() {
	var db = util.Mydb.Db
	err := db.AutoMigrate(&Article{})
	if err != nil {
		log.Panicln("Article表创建失败")
	}
}
