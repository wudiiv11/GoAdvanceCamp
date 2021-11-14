package usecase

import (
	"context"
	"geek.xingx/egPractice/internal/server/domain"
)

type article struct {
	repo domain.IArticleRepo
}

func NewArticleUseCase(repo domain.IArticleRepo) domain.IArticleUseCase {
	return &article{repo: repo}
}

func (u *article) CreateArticle(ctx context.Context, article *domain.Article) error {
	return u.repo.CreateArticle(ctx, article)
}
