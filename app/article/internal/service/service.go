package service

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	article "kratos-realworld/api/article/service/v1"
	"kratos-realworld/app/article/internal/biz"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewArticleService)

type ArticleService struct {
	article.UnimplementedArticleServer

	uc  *biz.ArticleUseCase
	log *log.Helper
}

func NewArticleService(uc *biz.ArticleUseCase, logger log.Logger) *ArticleService {
	return &ArticleService{
		uc:  uc,
		log: log.NewHelper(log.With(logger, "module", "service/article")),
	}
}
