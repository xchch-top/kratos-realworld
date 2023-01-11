package service

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	api "kratos-realworld/api/article/service/v1"
	"kratos-realworld/app/article/internal/biz"
	"kratos-realworld/pkg/middleware/auth"
)

func (s *ArticleService) GetArticle(ctx context.Context, req *api.GetArticleRequest) (*api.SingleArticleReply, error) {
	ba, err := s.uc.GetArticleBySlug(ctx, req.GetSlug())
	if err != nil {
		return nil, err
	}
	return &api.SingleArticleReply{
		Article: &api.ArticleDTO{
			Slug:        ba.Slug,
			Title:       ba.Title,
			Description: ba.Description,
			Body:        ba.Body,
		},
	}, nil
}

func (s *ArticleService) CreateArticle(ctx context.Context, req *api.CreateArticleRequest) (
	*api.SingleArticleReply, error) {
	authUser, err := auth.GetAuthUser(ctx)
	if err != nil {
		return nil, err
	}

	bizArticle := biz.Article{
		Title:       req.Article.Title,
		Description: req.Article.Description,
		Body:        req.Article.Body,
		AuthorID:    authUser.Id,
	}
	id, err := s.uc.CreateArticle(ctx, &bizArticle)
	if err != nil {
		return nil, err
	}

	ba, err := s.uc.GetArticle(ctx, id)
	if err != nil {
		return nil, err
	}
	return &api.SingleArticleReply{
		Article: &api.ArticleDTO{
			Slug:        ba.Slug,
			Title:       ba.Title,
			Description: ba.Description,
			Body:        ba.Body,
		},
	}, nil
}

func (s *ArticleService) UpdateArticle(ctx context.Context, req *api.UpdateArticleRequest) (*api.SingleArticleReply, error) {
	return nil, nil
}

func (s *ArticleService) DeleteArticle(ctx context.Context, req *api.DeleteArticleRequest) (*empty.Empty, error) {
	return nil, nil
}

func (s *ArticleService) AddComment(ctx context.Context, req *api.AddCommentRequest) (*api.SingleCommentReply, error) {
	return nil, nil
}

func (s *ArticleService) GetComments(ctx context.Context, req *api.GetCommentsRequest) (*api.MultipleCommentsReply, error) {
	return nil, nil
}

func (s *ArticleService) DeleteComment(ctx context.Context, req *api.DeleteCommentRequest) (*empty.Empty, error) {
	return nil, nil
}

func (s *ArticleService) FeedArticles(ctx context.Context, req *api.FeedArticlesRequest) (*api.MultipleArticlesReply, error) {
	return nil, nil
}

func (s *ArticleService) ListArticles(ctx context.Context, req *api.ListArticlesRequest) (*api.MultipleArticlesReply, error) {
	bas, err := s.uc.ListArticle(ctx)

	var articles []*api.ArticleDTO
	for _, ba := range bas {
		article := api.ArticleDTO{
			Slug:        ba.Slug,
			Title:       ba.Title,
			Description: ba.Description,
			Body:        ba.Body,
		}
		articles = append(articles, &article)
	}

	return &api.MultipleArticlesReply{
		Articles:      articles,
		ArticlesCount: int32(len(articles)),
	}, err
}

func (s *ArticleService) GetTags(ctx context.Context, req *empty.Empty) (*api.TagListReply, error) {
	return nil, nil
}

func (s *ArticleService) FavoriteArticle(ctx context.Context, req *api.FavoriteArticleRequest) (*api.SingleArticleReply, error) {
	return nil, nil
}

func (s *ArticleService) UnfavoriteArticle(ctx context.Context, req *api.UnfavoriteArticleRequest) (*empty.Empty, error) {
	return nil, nil
}
