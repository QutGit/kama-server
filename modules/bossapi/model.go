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

/** 创建实体 */
type Article struct {
	ID          string `gorm:"column:id" json:"id"`
	TermId      string `gorm:"column:term_id" json:"termId"`
	Title       string `gorm:"column:title" json:"title"`
	ImgUrl      string `gorm:"column:img_url" json:"imgUrl"`
	UserId      string `gorm:"column:user_id" json:"userId"`
	Description string `gorm:"column:description" json:"description"`
	CreateTime  string `gorm:"column:create_time" json:"createTime"`
	UpdateTime  string `gorm:"column:update_time" json:"updateTime"`
	Deleted     int    `gorm:"column:deleted" json:"deleted"`
}

/** 列表实体 */
type Articles struct {
	Id          string `gorm:"column:id" json:"id"`
	TermId      string `gorm:"column:term_id" json:"termId"`
	Title       string `gorm:"column:title" json:"title"`
	ImgUrl      string `gorm:"column:img_url" json:"imgUrl"`
	UserId      string `gorm:"column:user_id" json:"userId"`
	Description string `gorm:"column:description" json:"description"`
	CreateTime  string `gorm:"column:create_time" json:"createTime"`
	Deleted     int    `gorm:"column:deleted" json:"deleted"`
}

/**
  分页数据信息
*/
type ArticleEntity struct {
	Total int64      `json:"total"`
	List  []Articles `json:"list"`
}

/** 列表请求参数 */
type ArticleParam struct {
	UserId    string `gorm:"column:user_id" json:"userId"`
	Title     string `gorm:"column:title" json:"title"`
	StartTime string `json:"startTime"`
	EndTime   string `json:"endTime"`
	TermId    string `gorm:"column:term_id" json:"termId"`
	Deleted   string `gorm:"column:deleted" json:"deleted"`
	Limit     int64  `json:"limit"`
	Offset    int64  `json:"offset"`
}

func (m *category) String() string {
	return fmt.Sprintf("example.wpPosts<%s,%s>", m.TermId, m.Name)
}
