package example

import (
	"fmt"
	"golang-started/db"
)

// init 自动进行数据库初始化
func init() {
	err := db.GetDB().AutoMigrate(&wpPosts{})
	if err != nil {
		panic(err)
	}
}

// Courseware 课件模型
// type Example struct {
// 	util.CommonModel
// 	// Name 例子名称
// 	Name string `gorm:"column:name" json:"name"`
// }

type wpPosts struct {
	ID           string `gorm:"column:ID" json:"id"`
	Post_author  string `gorm:"column:post_author" json:"postAuthor"`
	Post_date    string `gorm:"column:post_date" json:"postDate"`
	Post_content string `gorm:"column:post_content" json:"postContent"`
}

func (m *wpPosts) String() string {
	return fmt.Sprintf("example.wpPosts<%s,%s>", m.ID, m.Post_content)
}
