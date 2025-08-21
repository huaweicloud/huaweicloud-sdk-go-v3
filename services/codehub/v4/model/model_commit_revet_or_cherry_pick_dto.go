package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CommitRevetOrCherryPickDto struct {

	// revert目标分支
	Branch string `json:"branch"`

	// 是否使用创建MR的形式Revert
	WithNewMergeRequest *bool `json:"with_new_merge_request,omitempty"`

	// 提交信息
	Message *string `json:"message,omitempty"`
}

func (o CommitRevetOrCherryPickDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CommitRevetOrCherryPickDto struct{}"
	}

	return strings.Join([]string{"CommitRevetOrCherryPickDto", string(data)}, " ")
}
