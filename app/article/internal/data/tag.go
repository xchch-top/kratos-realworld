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

	return &biz.Tag{
		Id:   tag.Id,
		Name: tag.Name,
	}, result.Error
}

func (t *tagRepo) GetTag(ctx context.Context, id uint64) (*biz.Tag, error) {
	var tag Tag
	tag.Id = id
	result := t.data.db.First(&tag)

	return &biz.Tag{
		Id:   tag.Id,
		Name: tag.Name,
	}, result.Error
}

func (t *tagRepo) CreateTag(ctx context.Context, tagName string) (id uint64, err error) {
	tag := &Tag{
		Name: tagName,
	}
	result := t.data.db.Create(tag)
	return tag.Id, result.Error
}

func (r *tagRepo) SaveArticleTagRelation(ctx context.Context, articleId uint64, tagIds []uint64) error {
	var ats []*ArticleTag
	for _, tagId := range tagIds {
		articleTag := ArticleTag{
			ArticleId: articleId,
			TagId:     tagId,
		}
		ats = append(ats, &articleTag)
	}
	return r.data.db.Table("article_tag").Create(&ats).Error
}

func (r *tagRepo) DeleteRelationByArticle(ctx context.Context, articleId uint64) error {
	result := r.data.db.Where("article_id = ?", articleId).Delete(&ArticleTag{})
	return result.Error
}

func (r *tagRepo) ListArticleByTagId(ctx context.Context, tagId uint64) ([]*biz.Article, error) {
	var articles []*Article
	result := r.data.db.Joins("left join article_tag on article.id = article_tag.article_id").Where("article_tag.tag_id = ?", tagId).Find(&articles)

	var bas []*biz.Article
	for _, a := range articles {
		ba := biz.Article{
			Id:          a.Id,
			Slug:        a.Slug,
			Title:       a.Title,
			Description: a.Description,
			Body:        a.Body,
			CreatedAt:   a.CreatedAt,
			UpdatedAt:   a.UpdatedAt,
			AuthorID:    a.AuthorID,
		}
		bas = append(bas, &ba)
	}

	return bas, result.Error
}
