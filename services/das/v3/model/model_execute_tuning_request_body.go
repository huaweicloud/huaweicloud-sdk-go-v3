package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExecuteTuningRequestBody 执行调优请求体
type ExecuteTuningRequestBody struct {

	// 数据库名称
	DatabaseName string `json:"database_name"`

	// schema名称，诊断实例类型为postgresql时可用
	SchemaName *string `json:"schema_name,omitempty"`

	// 诊断的SQL语句
	SqlScript string `json:"sql_script"`

	// 执行节点类型，取值范围：master（主节点）、slave（副节点）、readreplica（只读节点）
	NodeType *string `json:"node_type,omitempty"`

	// 执行节点ID，实例节点的唯一标识
	NodeId *string `json:"node_id,omitempty"`

	// 是否使用search_path作为环境变量，诊断实例类型为postgresql时可用
	UseDefaultSearchPath *bool `json:"use_default_search_path,omitempty"`
}

func (o ExecuteTuningRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecuteTuningRequestBody struct{}"
	}

	return strings.Join([]string{"ExecuteTuningRequestBody", string(data)}, " ")
}
