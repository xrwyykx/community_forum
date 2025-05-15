package dto

import (
	"communityforum/models"
	"time"
)

type GetPostMap struct {
	models.Page
	Content string `json:"content" gorm:"column:content"`
	Status  int    `json:"status" gorm:"column:status"` //0按标题或者内容  1按作者名字
}

type PostDetail struct {
	PostId   int    `json:"post_id" gorm:"column:post_id"`
	UserName string `json:"userName" gorm:"column:username"` //要注意gorm和前面用户寻找的结构体里面的gorm是否相同
	Title    string `json:"title" gorm:"column:title;not null"`
	Content  string `json:"content" gorm:"column:content;not null"`
	//CreateTime time.Time `json:"createTime" gorm:"column:createTime;not null"`
	CreateTime    time.Time `json:"-" gorm:"column:create_time;not null"`
	CreateTimeMar string    `json:"createTime"` //前端展现以及传过来的数据是json
	//IsTop      bool      `json:"is_top" gorm:"column:is_top;default:false"`
	//IsElite    bool      `json:"is_elite" gorm:"column:is_elite;default:false"`
}

type AddPostMap struct {
	//UserId  int64  `json:"userId" gorm:"column:userId"`
	Title   string `json:"title" gorm:"column:title"`
	Content string `json:"content" gorm:"column:content"`
}

type UpdatePostMap struct {
	PostId  int    `json:"post_id" gorm:"column:postId"` //那边过来是json字段，操作数据库是gorm
	Title   string `json:"title,omitempty" gorm:"column:title"`
	Content string `json:"content,omitempty" gorm:"column:content"`
}

type PostIdMap struct {
	PostId int `json:"post_id" gorm:"column:postId"`
} //但是如果重构成每个表的Id都没有带有表明前缀，就不需要单独加一个deletePostMap,每个删除的结构体都可以通用
//以后看用不用重构成软删除的

type ReplyPostMap struct {
	//回复可能回复的是帖子或者回复的是评论
	PostId   int    `json:"post_id" gorm:"column:postId"`
	ParentId int    `json:"parent_id" gorm:"column:parentId"` //为0表示回复的是帖子
	Content  string `json:"content" gorm:"column:content"`
}

type PostReply struct {
	PostId        int       `json:"post_id" gorm:"column:postId"`
	UserId        int64     `json:"user_id" gorm:"column:userId"`
	UserName      string    `json:"username" gorm:"column:username"`
	ParentId      int       `json:"parent_id" gorm:"parentId"`
	Content       string    `json:"content" gorm:"column:content"`
	CreateTime    time.Time `json:"-" gorm:"column:create_time;not null"`
	CreateTimeMar string    `json:"createTime"` //前端展现以及传过来的数据是json
}

type LikesPostMap struct {
}
