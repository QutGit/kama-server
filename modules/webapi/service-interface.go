package webapi

import "context"

type ServiceInterface interface {
	Get(ctx context.Context, id uint64) (context.Context, *wpPosts, error)
	Create(ctx context.Context, m *wpPosts) (context.Context, *wpPosts, error)
	GetCategory(ctx context.Context) (context.Context, []category, error)
}
