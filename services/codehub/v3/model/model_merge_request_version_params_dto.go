package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type MergeRequestVersionParamsDto struct {

	// MR最新变更id
	DiffId *int32 `json:"diff_id,omitempty"`

	// 目标分支最新提交
	StartSha *string `json:"start_sha,omitempty"`

	// 源分支最新提交
	CommitId *string `json:"commit_id,omitempty"`
}

func (o MergeRequestVersionParamsDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MergeRequestVersionParamsDto struct{}"
	}

	return strings.Join([]string{"MergeRequestVersionParamsDto", string(data)}, " ")
}
