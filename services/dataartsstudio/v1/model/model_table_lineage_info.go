package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TableLineageInfo 批量血缘信息
type TableLineageInfo struct {

	// 上游血缘资产guid
	SourceGuid *string `json:"source_guid,omitempty"`

	// 上游血缘资产Qname
	SourceQualifiedName *string `json:"source_qualified_name,omitempty"`

	// 上游血缘资产类型
	SourceType *string `json:"source_type,omitempty"`

	// 上游血缘资产数据库
	SourceDb *string `json:"source_db,omitempty"`

	// 上游血缘资产逻辑库
	SourceSchema *string `json:"source_schema,omitempty"`

	// 上游血缘资产表
	SourceTable *string `json:"source_table,omitempty"`

	// 下游血缘资产guid
	TargetGuid *string `json:"target_guid,omitempty"`

	// 下游血缘资产Qname
	TargetQualifiedName *string `json:"target_qualified_name,omitempty"`

	// 下游血缘资产类型
	TargetType *string `json:"target_type,omitempty"`

	// 下游血缘资产数据库
	TargetDb *string `json:"target_db,omitempty"`

	// 下游血缘资产逻辑库
	TargetSchema *string `json:"target_schema,omitempty"`

	// 下游血缘资产表
	TargetTable *string `json:"target_table,omitempty"`

	// 作业算子guid
	NodeGuid *string `json:"node_guid,omitempty"`

	// 作业算子名称
	NodeName *string `json:"node_name,omitempty"`

	// 作业算子类型
	NodeType *string `json:"node_type,omitempty"`

	// 作业算子Qname
	NodeQualifiedName *string `json:"node_qualified_name,omitempty"`

	// 作业算子类型所属空间
	WorkspaceId *string `json:"workspace_id,omitempty"`
}

func (o TableLineageInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TableLineageInfo struct{}"
	}

	return strings.Join([]string{"TableLineageInfo", string(data)}, " ")
}
