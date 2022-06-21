package db

import (
	"context"
	"time"

	"github.com/DontSerious/DouSheng/pkg/constants"
	"gorm.io/gorm"
)

// 数据库
type Video struct {
	gorm.Model
	AuthorId      int64 `gorm:"index"`
	PlayUrl       string
	CoverUrl      string
	FavoriteCount int64
	CommentCount  int64
	Title         string
}

// video表名
func (v *Video) TableName() string {
	return constants.VideoTableName
}

// 生成视频记录
func CreateVideo(ctx context.Context, v *Video) error {
	return MyDB.WithContext(ctx).Create(v).Error
}

// 查询视频列表
func QueryVideoList(ctx context.Context, vIds []int64) ([]*Video, error) {
	videoList := make([]*Video, 0)
	if err := MyDB.WithContext(ctx).Where("id = ?", vIds).Find(&videoList).Error; err != nil {
		return videoList, err
	}

	return videoList, nil
}

// 作者视频列表
func PublishList(ctx context.Context, authorId int64) ([]*Video, error) {
	videoList := make([]*Video, 0)
	if err := MyDB.WithContext(ctx).Where("author_id = ?", authorId).Error; err != nil {
		return nil, err
	}
	return videoList, nil
}

// 点赞操作 true为点赞 false为取消点赞
func Favorite(ctx context.Context, vId int64, actionType bool) error {
	tx := MyDB.WithContext(ctx)
	video := new(Video)
	video.ID = uint(vId)
	if actionType {
		if err := tx.Model(video).Update("favorite_count", gorm.Expr("favorite_count + ?", 1)).Error; err != nil {
			return err
		}
	} else {
		if err := tx.Model(video).Update("favorite_count", gorm.Expr("favorite_count - ?", 1)).Error; err != nil {
			return err
		}
	}
	return nil
}

// 评论操作 true为点赞 false为取消点赞
func Comment(ctx context.Context, vId int64, actionType bool) error {
	tx := MyDB.WithContext(ctx)
	video := new(Video)
	video.ID = uint(vId)
	if actionType {
		if err := tx.Model(video).Update("comment_count", gorm.Expr("comment_count + ?", 1)).Error; err != nil {
			return err
		}
	} else {
		if err := tx.Model(video).Update("comment_count", gorm.Expr("comment_count - ?", 1)).Error; err != nil {
			return err
		}
	}
	return nil
}

// 视频流
func Feed(ctx context.Context, lastTime time.Time) ([]*Video, int64, error) {
	videoList := make([]*Video, 0)
	var nextTime time.Time
	
	if err := MyDB.WithContext(ctx).Limit(30).Order("create_at DESC").Where("create_at < ?", lastTime).Find(&videoList).Error; err != nil {
		return nil, lastTime.Unix(), err
	}
	
	if len(videoList) == 0 {
		nextTime = lastTime
	} else {
		nextTime = videoList[len(videoList) - 1].CreatedAt
	}

	return videoList, nextTime.Unix(), nil
}
