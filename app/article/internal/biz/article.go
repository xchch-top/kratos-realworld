package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
)

type ArticleRepo interface {
	CreateArticle(ctx context.Context, article *Article) (uint64, error)
	GetArticle(ctx context.Context, id uint64) (*Article, error)
	GetArticleBySlug(ctx context.Context, slug string) (uint64, error)
	ListArticle(ctx context.Context) ([]*Article, error)
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

type Article struct {
	Id          uint64
	Slug        string
	Title       string `gorm:"size:200"`
	Description string `gorm:"size:200"`
	Body        string
	AuthorID    uint64
}

func (uc ArticleUseCase) CreateArticle(ctx context.Context, article *Article) (uint64, error) {
	id, err := uc.ar.CreateArticle(ctx, article)
	if err != nil {
		return 0, errors.InternalServer("create article", "创建文章失败")
	}
	return id, nil
}

func (uc ArticleUseCase) GetArticle(ctx context.Context, id uint64) (*Article, error) {
	article, err := uc.ar.GetArticle(ctx, id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.NotFound("article not found", "文章未找到")
	}
	if err != nil {
		return nil, err
	}
	return article, nil
}

func (uc ArticleUseCase) GetArticleBySlug(ctx context.Context, slug string) (*Article, error) {
	id, err := uc.ar.GetArticleBySlug(ctx, slug)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.NotFound("article not found", "文章未找到")
	}
	if err != nil {
		return nil, err
	}
	return uc.GetArticle(ctx, id)
}

func (uc ArticleUseCase) ListArticle(ctx context.Context) ([]*Article, error) {
	return uc.ar.ListArticle(ctx)
}
