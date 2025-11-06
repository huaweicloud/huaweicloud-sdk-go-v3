package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CommentPathDto 单文件下检视意见详情。
type CommentPathDto struct {

	// **参数解释：** 位于文件对比结果右侧的检视意见集合。
	New *[]LineDiscussionDto `json:"new,omitempty"`

	// **参数解释：** 位于文件对比结果左侧的检视意见集合。
	Old *[]LineDiscussionDto `json:"old,omitempty"`

	// **参数解释：** 文件名。
	Path *string `json:"path,omitempty"`
}

func (o CommentPathDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CommentPathDto struct{}"
	}

	return strings.Join([]string{"CommentPathDto", string(data)}, " ")
}
