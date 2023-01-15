package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"kratos-realworld/app/article/internal/biz"
)

type Tag struct {
	Id   uint64 `gorm:"primarykey"`
	Name string
}

func (t *Tag) TableName() string {
	return "tag"
}

type tagRepo struct {
	data *Data
	log  *log.Helper
}

func NewTagRepo(data *Data, logger log.Logger) biz.TagRepo {
	return &tagRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (t *tagRepo) GetTags(ctx context.Context) ([]*biz.Tag, error) {
	var tags []Tag
	result := t.data.db.Find(&tags)
	if result.Error != nil {
		return nil, result.Error
	}

	var bizTags []*biz.Tag
	for _, tag := range tags {
		bt := biz.Tag{
			Id:   tag.Id,
			Name: tag.Name,
		}
		bizTags = append(bizTags, &bt)
	}
	return bizTags, nil
}

func (t *tagRepo) GetByName(ctx context.Context, name string) (*biz.Tag, error) {
	var tag Tag
	tag.Name = name
	result := t.data.db.First(&tag)
	if result.Error != nil {
		return nil, result.Error
	}

	return &biz.Tag{
		Id:   tag.Id,
		Name: tag.Name,
	}, nil
}
