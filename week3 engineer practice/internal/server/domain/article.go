package domain

import "context"

type Article struct {
	Model
	Title   string `json:"title"`
	Content string `json:"content"`
}

type IArticleUseCase interface {
	CreateArticle(ctx context.Context, article *Article) error
}

type IArticleRepo interface {
	CreateArticle(ctx context.Context, article *Article) error
}
