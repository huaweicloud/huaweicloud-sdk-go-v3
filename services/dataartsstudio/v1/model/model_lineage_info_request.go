package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type LineageInfoRequest struct {

	// 集群id
	ClusterId *string `json:"cluster_id,omitempty"`

	// 数据连接类型
	DataSourceType *string `json:"data_source_type,omitempty"`

	// 数据连接id
	ConnectionId *string `json:"connection_id,omitempty"`

	// 数据连接名称
	ConnectionName *string `json:"connection_name,omitempty"`

	// 工作空间id
	WorkspaceId *string `json:"workspace_id,omitempty"`

	// 作业id
	JobId *string `json:"job_id,omitempty"`

	// 算子名称
	NodeName *string `json:"node_name,omitempty"`

	TableLineage *TableLineageV2 `json:"table_lineage,omitempty"`
}

func (o LineageInfoRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LineageInfoRequest struct{}"
	}

	return strings.Join([]string{"LineageInfoRequest", string(data)}, " ")
}
