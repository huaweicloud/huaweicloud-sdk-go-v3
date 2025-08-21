package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ReplyMergeRequestNoteDto 回复合并请求检视意见请求体
type ReplyMergeRequestNoteDto struct {

	// **参数解释：** 评论内容。
	Body *string `json:"body,omitempty"`
}

func (o ReplyMergeRequestNoteDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ReplyMergeRequestNoteDto struct{}"
	}

	return strings.Join([]string{"ReplyMergeRequestNoteDto", string(data)}, " ")
}
