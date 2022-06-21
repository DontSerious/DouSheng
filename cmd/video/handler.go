package main

import (
	"context"
	"github.com/DontSerious/DouSheng/kitex_gen/video"
)

// VideoServiceImpl implements the last service interface defined in the IDL.
type VideoServiceImpl struct{}

// Feed implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) Feed(ctx context.Context, req *video.FeedRequest) (resp *video.FeedResponse, err error) {
	// TODO: Your code here...
	return
}

// CreateVideo implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) CreateVideo(ctx context.Context, req *video.CreateVideoRequest) (resp *video.CreateVideoResponse, err error) {
	// TODO: Your code here...
	return
}

// PublishList implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) PublishList(ctx context.Context, req *video.PublishListRequest) (resp *video.PublishListResponse, err error) {
	// TODO: Your code here...
	return
}

// MGetVideo implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) MGetVideo(ctx context.Context, req *video.MGetVideoRequest) (resp *video.MGetVideoResponse, err error) {
	// TODO: Your code here...
	return
}

// Favorite implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) Favorite(ctx context.Context, req *video.FavoriteOperationRequest) (resp *video.FavoriteOperationResponse, err error) {
	// TODO: Your code here...
	return
}

// Comment implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) Comment(ctx context.Context, req *video.CommentOperationRequest) (resp *video.CommentOperationResponse, err error) {
	// TODO: Your code here...
	return
}
