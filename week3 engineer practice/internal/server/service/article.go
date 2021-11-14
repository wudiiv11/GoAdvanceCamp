package service

import (
	"context"

	v1 "geek.xingx/egPractice/api/product/app/v1"
	"geek.xingx/egPractice/internal/pkg/errcode"
	"geek.xingx/egPractice/internal/server/domain"

	"github.com/pkg/errors"
)

type ArticleService struct {
	useCase domain.IArticleUseCase
}

func NewArticleService(useCase domain.IArticleUseCase) *ArticleService {
	return &ArticleService{useCase: useCase}
}

func (u *ArticleService) CreateArticle(ctx context.Context, req *v1.CreateArticleReq) (*v1.CreateArticleResponse, error) {
	article := &domain.Article{
		Title:   req.Title,
		Content: req.Content,
	}

	err := u.useCase.CreateArticle(ctx, article)
	if err != nil {
		return nil, errors.Wrapf(errcode.ErrUnknown, "create article %s, %s", req.Title, err)
	}

	resp := &v1.CreateArticleResponse{}
	return resp, nil
}
