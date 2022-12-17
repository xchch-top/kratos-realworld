package service

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	realworld "kratos-realworld/api/realworld/v1"
)

func (s *RealworldService) GetProfile(ctx context.Context, req *realworld.GetProfileRequest) (*realworld.ProfileReply, error) {
	return nil, nil
}

func (s *RealworldService) FollowUser(ctx context.Context, req *realworld.FollowUserRequest) (*realworld.ProfileReply, error) {
	return nil, nil
}

func (s *RealworldService) UnfollowUser(ctx context.Context, req *realworld.UnfollowUserRequest) (*empty.Empty, error) {
	return nil, nil
}

func (s *RealworldService) GetArticle(ctx context.Context, req *realworld.GetArticleRequest) (*realworld.SingleArticleReply, error) {
	return nil, nil
}

func (s *RealworldService) CreateArticle(ctx context.Context, req *realworld.CreateArticleRequest) (*realworld.SingleArticleReply, error) {
	return nil, nil
}

func (s *RealworldService) UpdateArticle(ctx context.Context, req *realworld.UpdateArticleRequest) (*realworld.SingleArticleReply, error) {
	return nil, nil
}

func (s *RealworldService) DeleteArticle(ctx context.Context, req *realworld.DeleteArticleRequest) (*empty.Empty, error) {
	return nil, nil
}

func (s *RealworldService) AddComment(ctx context.Context, req *realworld.AddCommentRequest) (*realworld.SingleCommentReply, error) {
	return nil, nil
}

func (s *RealworldService) GetComments(ctx context.Context, req *realworld.GetCommentsRequest) (*realworld.MultipleCommentsReply, error) {
	return nil, nil
}

func (s *RealworldService) DeleteComment(ctx context.Context, req *realworld.DeleteCommentRequest) (*empty.Empty, error) {
	return nil, nil
}

func (s *RealworldService) FeedArticles(ctx context.Context, req *realworld.FeedArticlesRequest) (*realworld.MultipleArticlesReply, error) {
	return nil, nil
}

func (s *RealworldService) ListArticles(ctx context.Context, req *realworld.ListArticlesRequest) (*realworld.MultipleArticlesReply, error) {
	return nil, nil
}

func (s *RealworldService) GetTags(ctx context.Context, req *empty.Empty) (*realworld.TagListReply, error) {
	return nil, nil
}

func (s *RealworldService) FavoriteArticle(ctx context.Context, req *realworld.FavoriteArticleRequest) (*realworld.SingleArticleReply, error) {
	return nil, nil
}

func (s *RealworldService) UnfavoriteArticle(ctx context.Context, req *realworld.UnfavoriteArticleRequest) (*empty.Empty, error) {
	return nil, nil
}
