package biz

import "context"

type Tag struct {
	Id   uint64
	Name string
}

type TagRepo interface {
	GetTags(ctx context.Context) ([]*Tag, error)
	GetByName(ctx context.Context, name string) (*Tag, error)
}

func (uc *ArticleUseCase) GetTags(ctx context.Context) ([]*Tag, error) {
	return uc.tr.GetTags(ctx)
}
