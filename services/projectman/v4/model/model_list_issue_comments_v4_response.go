package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListIssueCommentsV4Response struct {
	// 评论总数

	Total *int32 `json:"total,omitempty"`
	// 品论列表

	Comments       *[]IssueCommentV4 `json:"comments,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o ListIssueCommentsV4Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListIssueCommentsV4Response struct{}"
	}

	return strings.Join([]string{"ListIssueCommentsV4Response", string(data)}, " ")
}
