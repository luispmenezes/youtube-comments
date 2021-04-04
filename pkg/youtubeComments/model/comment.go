package model

import "time"

type Comment struct {
	Id          string        `json:"id"`
	CommentText string        `json:"commentText"`
	LikeCount   int           `json:"likeCount"`
	IsHearted   bool          `json:"isHearted"`
	IsEdited    bool          `json:"isEdited"`
	Age         time.Duration `json:"age"`
	Author      Author        `json:"author"`
	Replies     []Comment     `json:"replies"`
}
