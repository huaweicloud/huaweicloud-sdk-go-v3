package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GradedCount 已保护已定级的资源统计
type GradedCount struct {

	// 总数
	Total int32 `json:"total"`

	// 合规资源数
	CompleteMatch *int32 `json:"complete_match,omitempty"`

	// 部分合规资源数
	PartialMatch *int32 `json:"partial_match,omitempty"`

	// 不合规资源数
	NoMatch *int32 `json:"no_match,omitempty"`
}

func (o GradedCount) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GradedCount struct{}"
	}

	return strings.Join([]string{"GradedCount", string(data)}, " ")
}
