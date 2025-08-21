package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DivergingCommitCounts struct {

	// 落后默认分支提交数量
	Behind *int32 `json:"behind,omitempty"`

	// 领先默认分支提交数量
	Ahead *int32 `json:"ahead,omitempty"`
}

func (o DivergingCommitCounts) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DivergingCommitCounts struct{}"
	}

	return strings.Join([]string{"DivergingCommitCounts", string(data)}, " ")
}
