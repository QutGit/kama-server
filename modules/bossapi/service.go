package bossapi

import (
	"context"
	"fmt"
	"golang-started/db"
)

// Service
type Service struct {
}

// Get 根据ID获取文章
func (that *Service) Get(ctx context.Context, id uint64) (context.Context, *category, error) {
	m := &category{}
	result := db.GetDB().Table("wp_posts").First(m, id)
	return ctx, m, result.Error
}

// Create 创建课件
func (that *Service) Create(ctx context.Context, m *category) (context.Context, *category, error) {
	result := db.GetDB().Create(m)
	return ctx, m, result.Error
}

// 获取分类目录
func (that *Service) GetCategory(ctx context.Context) (context.Context, []category, error) {
	var (
		cat = make([]category, 0)
	)
	_ = db.GetDB().Raw("SELECT wp_terms.term_id, wp_terms.name, wp_term_taxonomy.description, wp_term_taxonomy.icon_url, wp_terms.slug FROM wp_term_taxonomy JOIN wp_terms ON wp_terms.term_id=wp_term_taxonomy.term_id WHERE taxonomy='category'").Scan(&cat).Error
	return ctx, cat, nil
}

// 创建文章
func (that *Service) CreateArticle(ctx context.Context, arts []*Article) (context.Context, []*Article, error) {
	result := db.GetDB().Table("wp_list").Create(&arts)
	return ctx, arts, result.Error
}

// 获取文章列表
func (that *Service) GetArticles(ctx context.Context, param *ArticleParam) (context.Context, *ArticleEntity, error) {
	var (
		arts = make([]Articles, 0)
	)
	var count int64
	list := db.GetDB().Table("wp_list").Limit(param.Limit).Offset(param.Offset).Where(map[string]interface{}{"user_id": param.UserId}).Find(&arts)
	total := db.GetDB().Table("wp_list").Count(&count)
	fmt.Println("**************************")
	fmt.Println(&total)
	fmt.Println(&list)
	fmt.Println("**************************")
	// result := ArticleEntity{total, arts}
	return ctx, nil, list.Error
}
