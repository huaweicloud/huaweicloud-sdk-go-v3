package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ProcessStats 会话统计信息
type ProcessStats struct {

	// 统计项名称
	Key *string `json:"key,omitempty"`

	// 统计项值
	Value *string `json:"value,omitempty"`

	// 活跃数
	ActiveCount *int64 `json:"active_count,omitempty"`

	// 总数
	TotalCount *int64 `json:"total_count,omitempty"`
}

func (o ProcessStats) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProcessStats struct{}"
	}

	return strings.Join([]string{"ProcessStats", string(data)}, " ")
}
