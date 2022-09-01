package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type IssueCommentV4 struct {

	// 评论内容
	Comment *string `json:"comment,omitempty" xml:"comment"`

	// 评论id
	Id *int32 `json:"id,omitempty" xml:"id"`

	// 评论时间
	CreatedTime *string `json:"created_time,omitempty" xml:"created_time"`

	User *CommentUserV4 `json:"user,omitempty" xml:"user"`
}

func (o IssueCommentV4) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IssueCommentV4 struct{}"
	}

	return strings.Join([]string{"IssueCommentV4", string(data)}, " ")
}
