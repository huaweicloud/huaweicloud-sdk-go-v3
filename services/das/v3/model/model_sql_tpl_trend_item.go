package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SqlTplTrendItem SQL模板趋势项
type SqlTplTrendItem struct {

	// 执行时间 ms
	ExecuteAt *int64 `json:"execute_at,omitempty"`

	// 耗时在500ms的个数
	QueryTimeIn500ms *int64 `json:"query_time_in500ms,omitempty"`

	// 耗时在100ms的个数
	QueryTimeIn100ms *int64 `json:"query_time_in100ms,omitempty"`

	// 耗时在1s的个数
	QueryTimeIn1s *int64 `json:"query_time_in1s,omitempty"`

	// 耗时超过1s的个数
	QueryTimeOver1s *int64 `json:"query_time_over1s,omitempty"`

	// 总个数
	QueryExecutions *int64 `json:"query_executions,omitempty"`
}

func (o SqlTplTrendItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SqlTplTrendItem struct{}"
	}

	return strings.Join([]string{"SqlTplTrendItem", string(data)}, " ")
}
