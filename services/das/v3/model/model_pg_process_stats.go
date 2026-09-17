package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PgProcessStats PgProcessStats
type PgProcessStats struct {

	// 参数名
	Key *string `json:"key,omitempty"`

	// 参数值
	Value *string `json:"value,omitempty"`

	// 活跃数
	ActiveCount *int64 `json:"active_count,omitempty"`

	// 总数
	TotalCount *int64 `json:"total_count,omitempty"`
}

func (o PgProcessStats) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PgProcessStats struct{}"
	}

	return strings.Join([]string{"PgProcessStats", string(data)}, " ")
}
