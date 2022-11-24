package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DivergingCommitCounts struct {

	// 领先提交数
	Ahead *float64 `json:"ahead,omitempty"`

	// 滞后提交数
	Behind *float64 `json:"behind,omitempty"`
}

func (o DivergingCommitCounts) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DivergingCommitCounts struct{}"
	}

	return strings.Join([]string{"DivergingCommitCounts", string(data)}, " ")
}
