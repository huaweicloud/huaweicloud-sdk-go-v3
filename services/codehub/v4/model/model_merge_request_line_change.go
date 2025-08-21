package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type MergeRequestLineChange struct {

	// 合并请求新增代码行数
	AddedLines *int32 `json:"added_lines,omitempty"`

	// 合并请求删除代码行数
	RemovedLines *int32 `json:"removed_lines,omitempty"`
}

func (o MergeRequestLineChange) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MergeRequestLineChange struct{}"
	}

	return strings.Join([]string{"MergeRequestLineChange", string(data)}, " ")
}
