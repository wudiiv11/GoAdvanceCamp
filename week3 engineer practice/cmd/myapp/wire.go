//go:build wireinject

package main

import (
	"geek.xingx/egPractice/internal/server/repo"
	"geek.xingx/egPractice/internal/server/service"
	"geek.xingx/egPractice/internal/server/usecase"
	"github.com/google/wire"
)

var (
	articleSet = wire.NewSet(
		usecase.NewArticleUseCase,
		repo.NewArticleRepository,
		service.NewArticleService,
	)
)

func InitArticleService() (*service.ArticleService, error) {
	wire.Build(articleSet)
	return new(service.ArticleService), nil
}
