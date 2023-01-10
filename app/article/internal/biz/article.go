package biz

import "github.com/go-kratos/kratos/v2/log"

type ArticleRepo interface {
}

type CommentRepo interface {
}

type TagRepo interface {
}

type ArticleUseCase struct {
	ar ArticleRepo
	cr CommentRepo
	tr TagRepo

	log *log.Helper
}

func NewArticleUseCase(ar ArticleRepo, cr CommentRepo, tr TagRepo, logger log.Logger) *ArticleUseCase {
	return &ArticleUseCase{
		ar:  ar,
		cr:  cr,
		tr:  tr,
		log: log.NewHelper(logger),
	}
}
