package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CommentData 用户提交的评论信息。
type CommentData struct {

	// 标题。
	CommentTitle *string `json:"comment_title,omitempty"`

	// 消息。
	CommentMessage *string `json:"comment_message,omitempty"`

	// 附件名字
	AttachmentName *[]string `json:"attachment_name,omitempty"`
}

func (o CommentData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CommentData struct{}"
	}

	return strings.Join([]string{"CommentData", string(data)}, " ")
}
