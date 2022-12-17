package biz

import "github.com/go-kratos/kratos/v2/log"

type ArticleRepo interface {
}

type CommentRepo interface {
}

type TagRepo interface {
}

type SocialUseCase struct {
	ar ArticleRepo
	cr CommentRepo
	tr TagRepo

	log *log.Helper
}

func NewSocialUseCase(ar ArticleRepo, cr CommentRepo, tr TagRepo, logger log.Logger) *SocialUseCase {
	return &SocialUseCase{
		ar:  ar,
		cr:  cr,
		tr:  tr,
		log: log.NewHelper(logger),
	}
}
