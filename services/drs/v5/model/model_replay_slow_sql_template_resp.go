package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ReplaySlowSqlTemplateResp 慢SQL统计信息数据项
type ReplaySlowSqlTemplateResp struct {

	// SQL语句模板
	SqlTemplate string `json:"sql_template"`

	// SQL语句模板MD5
	SqlTemplateMd5 *string `json:"sql_template_md5,omitempty"`

	// 目标库别名
	TargetName *string `json:"target_name,omitempty"`

	// schema名称
	SchemaName *string `json:"schema_name,omitempty"`

	// SQL类型
	QueryType *string `json:"query_type,omitempty"`

	// SQL执行最小耗时
	MinLatency *int64 `json:"min_latency,omitempty"`

	// SQL执行最大耗时
	MaxLatency *int64 `json:"max_latency,omitempty"`

	// SQL执行平均耗时
	AvgLatency int64 `json:"avg_latency"`

	// SQL执行总的耗时
	TotalLatency *int64 `json:"total_latency,omitempty"`

	// 目标库类型
	TargetType *string `json:"target_type,omitempty"`

	// SQL数量
	Count int64 `json:"count"`
}

func (o ReplaySlowSqlTemplateResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ReplaySlowSqlTemplateResp struct{}"
	}

	return strings.Join([]string{"ReplaySlowSqlTemplateResp", string(data)}, " ")
}
