package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExTimeTrendItem 执行时间趋势项
type ExTimeTrendItem struct {

	// SQL执行的时间点
	ExecuteAt *int64 `json:"execute_at,omitempty"`

	// SQL执行耗时
	ExecuteTime *float64 `json:"execute_time,omitempty"`
}

func (o ExTimeTrendItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExTimeTrendItem struct{}"
	}

	return strings.Join([]string{"ExTimeTrendItem", string(data)}, " ")
}
