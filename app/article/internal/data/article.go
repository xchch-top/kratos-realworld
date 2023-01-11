package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/uuid"
	"kratos-realworld/app/article/internal/biz"
)

type Article struct {
	Model
	Slug        string `gorm:"size:200"`
	Title       string `gorm:"size:200"`
	Description string `gorm:"size:200"`
	Body        string
	Tags        []Tag `gorm:"many2many:article_tags;"`
	AuthorID    uint64
	// Author         User
	FavoritesCount uint32
}

type articleRepo struct {
	data *Data
	log  *log.Helper
}

func NewArticleRepo(data *Data, logger log.Logger) biz.ArticleRepo {
	return &articleRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *articleRepo) CreateArticle(ctx context.Context, ba *biz.Article) (uint64, error) {
	article := Article{
		Slug:        uuid.New().String(),
		Title:       ba.Title,
		Description: ba.Description,
		Body:        ba.Body,
		AuthorID:    ba.AuthorID,
	}
	result := r.data.db.Create(&article)
	return article.Id, result.Error
}

func (r *articleRepo) GetArticle(ctx context.Context, id uint64) (*biz.Article, error) {
	var article Article
	article.Id = id
	result := r.data.db.First(&article)

	return &biz.Article{
		Id:          article.Id,
		Slug:        article.Slug,
		Title:       article.Title,
		Description: article.Description,
		Body:        article.Body,
	}, result.Error
}

func (r *articleRepo) GetArticleBySlug(ctx context.Context, slug string) (uint64, error) {
	var article Article
	result := r.data.db.Select("id").Where("slug = ?", slug).First(&article)
	return article.Id, result.Error
}

func (r *articleRepo) ListArticle(ctx context.Context) ([]*biz.Article, error) {
	var articles []*Article
	result := r.data.db.Find(&articles)

	var bas []*biz.Article
	for _, a := range articles {
		ba := biz.Article{
			Id:          a.Id,
			Slug:        a.Slug,
			Title:       a.Title,
			Description: a.Description,
			Body:        a.Body,
		}
		bas = append(bas, &ba)
	}

	return bas, result.Error
}
