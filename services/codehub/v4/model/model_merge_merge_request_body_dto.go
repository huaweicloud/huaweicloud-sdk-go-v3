package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type MergeMergeRequestBodyDto struct {

	// 压缩合并信息
	SquashCommitMessage *string `json:"squash_commit_message,omitempty"`

	// 压缩合并
	Squash *bool `json:"squash,omitempty"`

	// 强制合并
	ForceMerge *bool `json:"force_merge,omitempty"`
}

func (o MergeMergeRequestBodyDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MergeMergeRequestBodyDto struct{}"
	}

	return strings.Join([]string{"MergeMergeRequestBodyDto", string(data)}, " ")
}
