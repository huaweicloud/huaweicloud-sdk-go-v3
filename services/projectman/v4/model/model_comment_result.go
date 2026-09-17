package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CommentResult IPD工作项评论列表查询结果
type CommentResult struct {

	// **参数解释**： 符合过滤条件的工作项评论总数。 **取值范围**： 不涉及。
	Total *int32 `json:"total,omitempty"`

	// **参数解释**： 工作项评论列表。 **取值范围**： 不涉及。
	CommentList *[]CommentEntity `json:"comment_list,omitempty"`
}

func (o CommentResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CommentResult struct{}"
	}

	return strings.Join([]string{"CommentResult", string(data)}, " ")
}
