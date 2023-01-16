package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"gorm.io/gorm"
	userApi "kratos-realworld/api/user/service/v1"
	"time"
)

type ArticleRepo interface {
	CreateArticle(ctx context.Context, article *Article) (uint64, error)
	GetArticle(ctx context.Context, id uint64) (*Article, error)
	GetArticleBySlug(ctx context.Context, slug string) (uint64, error)
	ListArticle(ctx context.Context, listParam *ListParam) ([]*Article, error)
}

type CommentRepo interface {
}

type ArticleUseCase struct {
	ar ArticleRepo
	cr CommentRepo
	tr TagRepo
	uc userApi.UserClient

	log *log.Helper
}

func NewArticleUseCase(ar ArticleRepo, cr CommentRepo, tr TagRepo, uc userApi.UserClient, logger log.Logger) *ArticleUseCase {
	return &ArticleUseCase{
		ar:  ar,
		cr:  cr,
		tr:  tr,
		uc:  uc,
		log: log.NewHelper(logger),
	}
}

func NewUserClient(r registry.Discovery) userApi.UserClient {
	conn, err := grpc.DialInsecure(context.Background(),
		grpc.WithEndpoint("discovery:///realworld.api.user"),
		grpc.WithDiscovery(r))
	if err != nil {
		panic(err)
	}
	userClient := userApi.NewUserClient(conn)
	return userClient

}

type Article struct {
	Id          uint64
	Slug        string
	Title       string
	Description string
	Body        string
	AuthorID    uint64
	Author      *Author
	CreatedAt   time.Time
	UpdatedAt   time.Time
	TagList     []string
}

type Author struct {
	AuthorID  uint64
	Username  string
	Bio       string
	Image     string
	Following bool
}

func (uc *ArticleUseCase) CreateArticle(ctx context.Context, article *Article) (*Article, error) {
	id, err := uc.ar.CreateArticle(ctx, article)
	if err != nil {
		return nil, errors.InternalServer("create article", "创建文章失败")
	}

	err = uc.SaveTags(ctx, id, article.TagList)
	if err != nil {
		return nil, errors.InternalServer("create article", "创建文章失败")
	}

	return uc.GetArticle(ctx, id)
}

func (uc *ArticleUseCase) GetArticle(ctx context.Context, id uint64) (*Article, error) {
	article, err := uc.ar.GetArticle(ctx, id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.NotFound("article not found", "文章未找到")
	}
	if err != nil {
		return nil, err
	}

	author, err := uc.uc.GetUserById(ctx, &userApi.GetUserByIdRequest{Id: article.AuthorID})
	if err != nil {
		return nil, errors.InternalServer("author not found", "文章作者不存在")
	}
	article.Author = &Author{
		AuthorID:  article.AuthorID,
		Username:  author.User.Username,
		Bio:       author.User.Bio,
		Image:     author.User.Image,
		Following: false,
	}

	var tagList []string
	tagList, err = uc.GetTagIdsByArticle(ctx, id)
	if err != nil {
		return nil, err
	}
	article.TagList = tagList

	return article, nil
}

func (uc *ArticleUseCase) GetArticleBySlug(ctx context.Context, slug string) (*Article, error) {
	id, err := uc.ar.GetArticleBySlug(ctx, slug)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.NotFound("article not found", "文章未找到")
	}
	if err != nil {
		return nil, err
	}
	return uc.GetArticle(ctx, id)
}

type ListParam struct {
	Tag string
}

func (uc *ArticleUseCase) ListArticle(ctx context.Context, listParam *ListParam) ([]*Article, error) {
	var articles []*Article
	var err error
	if listParam.Tag != "" {
		bizTag, err := uc.tr.GetByName(ctx, listParam.Tag)
		if err != nil {
			return nil, err
		}
		articles, err = uc.tr.ListArticleByTagId(ctx, bizTag.Id)
		if err != nil {
			return nil, err
		}
	} else {
		articles, err = uc.ar.ListArticle(ctx, listParam)
		if err != nil {
			return nil, err
		}
	}

	for _, article := range articles {
		author, err := uc.uc.GetUserById(ctx, &userApi.GetUserByIdRequest{Id: article.AuthorID})
		if err != nil {
			log.Errorf("auth %s not found", article.AuthorID)
		} else {
			article.Author = &Author{
				AuthorID: article.AuthorID,
				Username: author.User.Username,
				Bio:      author.User.Bio,
				Image:    author.User.Image,
			}
		}
	}

	return articles, err
}
