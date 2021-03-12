package bossapi

import (
	"fmt"
	"golang-started/db"
)

// init 自动进行数据库初始化
func init() {
	err := db.GetDB().AutoMigrate(&category{})
	if err != nil {
		panic(err)
	}
}

type category struct {
	TermId      string `gorm:"column:term_id" json:"termId"`
	Name        string `gorm:"column:name" json:"name"`
	Description string `gorm:"column:description" json:"description"`
	IconUrl     string `gorm:"column:icon_url;" json:"iconUrl"`
	Slug        string `gorm:"column:slug" json:"slug"`
}

type Article struct {
	TermId string `gorm:"column:term_id" json:"termId"`
	Title  string `gorm:"column:title" json:"title"`
	ImgUrl string `gorm:"column:img_url" json:"imgUrl"`
	UserId string `gorm:"column:user_id" json:"userId"`
}

type Articles struct {
	Id     string `gorm:"column:id" json:"id"`
	TermId string `gorm:"column:term_id" json:"termId"`
	Title  string `gorm:"column:title" json:"title"`
	ImgUrl string `gorm:"column:img_url" json:"imgUrl"`
	UserId string `gorm:"column:user_id" json:"userId"`
}

type ArticleEntity struct {
	Total int      `json:"total"`
	List  Articles `json:"list"`
}

type ArticleParam struct {
	UserId string `gorm:"column:user_id" json:"userId"`
	Limit  int    `json:"limit"`
	Offset int    `json:"offset"`
}

func (m *category) String() string {
	return fmt.Sprintf("example.wpPosts<%s,%s>", m.TermId, m.Name)
}
