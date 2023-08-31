package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateTuningReq struct {

	// 数据库名称
	DatabaseName string `json:"database_name"`

	// schema名称
	SchemaName *string `json:"schema_name,omitempty"`

	// sql脚本
	SqlScript string `json:"sql_script"`

	// 节点类型
	NodeType *string `json:"node_type,omitempty"`

	// 节点Id
	NodeId *string `json:"node_id,omitempty"`
}

func (o CreateTuningReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTuningReq struct{}"
	}

	return strings.Join([]string{"CreateTuningReq", string(data)}, " ")
}
