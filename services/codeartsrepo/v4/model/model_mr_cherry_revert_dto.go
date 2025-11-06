package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type MrCherryRevertDto struct {

	// 分支名
	Branch string `json:"branch"`

	// 是否通过新建合并请求来Cherry pick或revert合并请求
	WithNewMergeRequest *bool `json:"with_new_merge_request,omitempty"`

	// 提交信息
	Message *string `json:"message,omitempty"`
}

func (o MrCherryRevertDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MrCherryRevertDto struct{}"
	}

	return strings.Join([]string{"MrCherryRevertDto", string(data)}, " ")
}
