package model

func InitDb() {
	user := User{}
	article := Article{}
	user.Generate()
	article.Generate()
}
