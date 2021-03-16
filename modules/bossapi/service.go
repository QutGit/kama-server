package bossapi

import (
	"context"
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
	Db := db.GetDB().Table("wp_list")
	if param.UserId != "" {
		Db.Where(map[string]interface{}{"user_id": param.UserId})
	}
	if param.Deleted != "" {
		Db.Where(map[string]interface{}{"deleted": param.Deleted})
	}
	if param.TermId != "" {
		Db.Where(map[string]interface{}{"term_id": param.TermId})
	}
	if param.StartTime != "" && param.EndTime != "" {
		Db.Where("create_time BETWEEN ? AND ?", param.StartTime, param.EndTime)
	}
	if param.Title != "" {
		Db.Where("title LIKE ?", "%"+param.Title+"%")
	}
	_ = Db.Count(&count)
	_ = Db.Limit(int(param.Limit)).Offset(int(param.Offset)).Order("create_time desc").Find(&arts)
	return ctx, &ArticleEntity{
		Total: count,
		List:  arts,
	}, nil
}

// 删除文章
func (that *Service) DeleteArticle(ctx context.Context, id string) (context.Context, error) {
	relult := db.GetDB().Table("wp_list").Where(map[string]interface{}{"id": id}).Update("deleted", 1)
	return ctx, relult.Error
}

// 恢复文章
func (that *Service) RecoverArticle(ctx context.Context, id string) (context.Context, error) {
	relult := db.GetDB().Table("wp_list").Where(map[string]interface{}{"id": id}).Update("deleted", 0)
	return ctx, relult.Error
}

// 修改文章
func (that *Service) UpdateArticle(ctx context.Context, id string, termId string, title string, description string) (context.Context, error) {
	relult := db.GetDB().Table("wp_list").Where(map[string]interface{}{"id": id}).Updates(map[string]interface{}{"term_id": termId, "title": title, "description": description})
	return ctx, relult.Error
}
