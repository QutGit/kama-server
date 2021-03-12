package bossapi

import "context"

type ServiceInterface interface {
	Get(ctx context.Context, id uint64) (context.Context, *category, error)
	Create(ctx context.Context, m *category) (context.Context, *category, error)
	GetCategory(ctx context.Context) (context.Context, []category, error)
	CreateArticle(ctx context.Context, a []*Article) (context.Context, []*Article, error)
	GetArticles(ctx context.Context, param *ArticleParam) (context.Context, *ArticleEntity, error)
}
