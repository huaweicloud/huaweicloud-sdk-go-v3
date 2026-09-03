package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowExecutionPlanRequestBody 获取执行计划请求体
type ShowExecutionPlanRequestBody struct {

	// 数据库名称
	DatabaseName string `json:"database_name"`

	// schema名称
	SchemaName *string `json:"schema_name,omitempty"`

	// SQL脚本
	SqlScript string `json:"sql_script"`

	// 实例节点ID，实例节点的唯一标识
	NodeId *string `json:"node_id,omitempty"`

	// 节点类型（master：主节点，slave：副节点，readreplica：只读节点）
	NodeType *string `json:"node_type,omitempty"`

	// PostgreSQL是否使用默认searchPath（仅在实例是PostgreSQL时可用）
	UseDefaultSearchPath *bool `json:"use_default_search_path,omitempty"`

	// 是否忽略限制
	IgnoreLimit *bool `json:"ignore_limit,omitempty"`

	// 每页记录数，取值范围：[0, 100]
	Perpage int32 `json:"perpage"`

	// 页码，取值范围：[0, 2^31-1]
	Curpage int32 `json:"curpage"`
}

func (o ShowExecutionPlanRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowExecutionPlanRequestBody struct{}"
	}

	return strings.Join([]string{"ShowExecutionPlanRequestBody", string(data)}, " ")
}
