package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"gorm.io/gorm"
)

type Tag struct {
	Id   uint64
	Name string
}

type TagRepo interface {
	GetTags(ctx context.Context) ([]*Tag, error)
	GetTag(ctx context.Context, id uint64) (*Tag, error)
	GetByName(ctx context.Context, name string) (*Tag, error)
	CreateTag(ctx context.Context, tagName string) (id uint64, err error)
	ListArticleByTagId(ctx context.Context, tagId uint64) ([]*Article, error)
	DeleteRelationByArticle(ctx context.Context, id uint64) error
	SaveArticleTagRelation(ctx context.Context, articleId uint64, tagIds []uint64) error
}

func (uc *ArticleUseCase) GetTags(ctx context.Context) ([]*Tag, error) {
	return uc.tr.GetTags(ctx)
}

func (uc *ArticleUseCase) SaveTags(ctx context.Context, articleId uint64, tagList []string) error {
	err := uc.tr.DeleteRelationByArticle(ctx, articleId)
	if err != nil {
		return err
	}

	var tagIds []uint64
	for _, tagName := range tagList {
		bizTag, err := uc.tr.GetByName(ctx, tagName)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			tagId, err := uc.tr.CreateTag(ctx, tagName)
			if err != nil {
				return err
			}

			bizTag, err = uc.tr.GetTag(ctx, tagId)
		}
		tagIds = append(tagIds, bizTag.Id)
	}

	return uc.tr.SaveArticleTagRelation(ctx, articleId, tagIds)
}

func (uc *ArticleUseCase) GetTagIdsByArticle(ctx context.Context, articleId uint64) ([]string, error) {
	return []string{"tag1", "tag2"}, nil
}
