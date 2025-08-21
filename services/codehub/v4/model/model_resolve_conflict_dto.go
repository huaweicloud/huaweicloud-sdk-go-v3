package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ResolveConflictDto struct {

	// **参数解释：** 提交信息
	CommitMessage string `json:"commit_message"`

	// **参数解释：** 冲突的文件
	Files []ResolveConflictFileDto `json:"files"`
}

func (o ResolveConflictDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResolveConflictDto struct{}"
	}

	return strings.Join([]string{"ResolveConflictDto", string(data)}, " ")
}
