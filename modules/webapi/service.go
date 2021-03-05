package webapi

import (
	"context"
	"golang-started/db"
)

// Service
type Service struct {
}

// Get 根据ID获取文章
func (that *Service) Get(ctx context.Context, id uint64) (context.Context, *wpPosts, error) {
	m := &wpPosts{}
	result := db.GetDB().Table("wp_posts").First(m, id)
	return ctx, m, result.Error
}

// Create 创建课件
func (that *Service) Create(ctx context.Context, m *wpPosts) (context.Context, *wpPosts, error) {
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
