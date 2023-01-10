package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"kratos-realworld/app/article/internal/biz"
)

type Article struct {
	Model
	Slug        string `gorm:"size:200"`
	Title       string `gorm:"size:200"`
	Description string `gorm:"size:200"`
	Body        string
	Tags        []Tag `gorm:"many2many:article_tags;"`
	AuthorID    uint
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

type ArticleCase struct {
	ar articleRepo
	cr commentRepo
	tr tagRepo

	log *log.Helper
}

func NewArticleCase(ar articleRepo, cr commentRepo, tr tagRepo, logger log.Logger) *ArticleCase {
	return &ArticleCase{
		ar:  ar,
		cr:  cr,
		tr:  tr,
		log: log.NewHelper(logger),
	}
}

func (s *ArticleCase) CreateArticle(ctx context.Context) error {
	return nil
}
