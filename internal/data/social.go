package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type articleRepo struct {
}

type commentRepo struct {
}

type tagRepo struct {
}

type socialUseCase struct {
	ar articleRepo
	cr commentRepo
	tr tagRepo

	log *log.Helper
}

func NewSocialUseCase(ar articleRepo, cr commentRepo, tr tagRepo, logger log.Logger) *socialUseCase {
	return &socialUseCase{
		ar:  ar,
		cr:  cr,
		tr:  tr,
		log: log.NewHelper(logger),
	}
}

func (s *socialUseCase) CreateArticle(ctx context.Context) error {
	return nil
}
