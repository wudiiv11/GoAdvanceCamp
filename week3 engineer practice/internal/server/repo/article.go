package repo

import (
	"context"
	"geek.xingx/egPractice/internal/server/domain"
)

type articleRepository struct {
}

func NewArticleRepository() domain.IArticleRepo {
	return &articleRepository{}
}

func (r *articleRepository) CreateArticle(ctx context.Context, article *domain.Article) error {
	return nil
}
