package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type MultiTaskMappingCreateBody struct {
	ExtInfo *MultiTaskInitElementExtInfo `json:"ext_info,omitempty"`

	// 源端数据源ID
	SourceDatasourceId string `json:"source_datasource_id"`

	// 目标端数据源ID
	TargetDatasourceId string `json:"target_datasource_id"`

	// 源端字段列表
	SourceColumns *[]MultiTaskColumnInfo `json:"source_columns,omitempty"`

	// 目标端字段列表
	TargetColumns *[]MultiTaskColumnInfo `json:"target_columns,omitempty"`

	// 源表名
	SourceTable *string `json:"source_table,omitempty"`

	// 目标表名
	TargetTable *string `json:"target_table,omitempty"`

	// 字段映射列表
	MappingColumns *[]MappingInfo `json:"mapping_columns,omitempty"`
}

func (o MultiTaskMappingCreateBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MultiTaskMappingCreateBody struct{}"
	}

	return strings.Join([]string{"MultiTaskMappingCreateBody", string(data)}, " ")
}
