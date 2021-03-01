package example

import (
	"context"
	"golang-started/db"
)

// Service 课件服务
type Service struct {
}

// Get 根据ID获取课件
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
