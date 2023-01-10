package service

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	articleSvc "kratos-realworld/api/article/service/v1"
)

func (s *ArticleService) GetArticle(ctx context.Context, req *articleSvc.GetArticleRequest) (*articleSvc.SingleArticleReply, error) {
	return nil, nil
}

func (s *ArticleService) CreateArticle(ctx context.Context, req *articleSvc.CreateArticleRequest) (
	*articleSvc.SingleArticleReply, error) {

	// s.sc.CreateArticle(ctx, biz.Create)

	return nil, nil
}

func (s *ArticleService) UpdateArticle(ctx context.Context, req *articleSvc.UpdateArticleRequest) (*articleSvc.SingleArticleReply, error) {
	return nil, nil
}

func (s *ArticleService) DeleteArticle(ctx context.Context, req *articleSvc.DeleteArticleRequest) (*empty.Empty, error) {
	return nil, nil
}

func (s *ArticleService) AddComment(ctx context.Context, req *articleSvc.AddCommentRequest) (*articleSvc.SingleCommentReply, error) {
	return nil, nil
}

func (s *ArticleService) GetComments(ctx context.Context, req *articleSvc.GetCommentsRequest) (*articleSvc.MultipleCommentsReply, error) {
	return nil, nil
}

func (s *ArticleService) DeleteComment(ctx context.Context, req *articleSvc.DeleteCommentRequest) (*empty.Empty, error) {
	return nil, nil
}

func (s *ArticleService) FeedArticles(ctx context.Context, req *articleSvc.FeedArticlesRequest) (*articleSvc.MultipleArticlesReply, error) {
	return nil, nil
}

func (s *ArticleService) ListArticles(ctx context.Context, req *articleSvc.ListArticlesRequest) (*articleSvc.MultipleArticlesReply, error) {
	return nil, nil
}

func (s *ArticleService) GetTags(ctx context.Context, req *empty.Empty) (*articleSvc.TagListReply, error) {
	return nil, nil
}

func (s *ArticleService) FavoriteArticle(ctx context.Context, req *articleSvc.FavoriteArticleRequest) (*articleSvc.SingleArticleReply, error) {
	return nil, nil
}

func (s *ArticleService) UnfavoriteArticle(ctx context.Context, req *articleSvc.UnfavoriteArticleRequest) (*empty.Empty, error) {
	return nil, nil
}
