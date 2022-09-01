package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type MultiTaskMappingCreateBody struct {
	ExtInfo *MultiTaskInitElementExtInfo `json:"ext_info,omitempty" xml:"ext_info"`

	// 源端数据源ID
	SourceDatasourceId *string `json:"source_datasource_id,omitempty" xml:"source_datasource_id"`

	// 目标端数据源ID
	TargetDatasourceId *string `json:"target_datasource_id,omitempty" xml:"target_datasource_id"`

	// 源端字段列表
	SourceColumns *[]MultiTaskColumnInfo `json:"source_columns,omitempty" xml:"source_columns"`

	// 目标端字段列表
	TargetColumns *[]MultiTaskColumnInfo `json:"target_columns,omitempty" xml:"target_columns"`

	// 源表名
	SourceTable *string `json:"source_table,omitempty" xml:"source_table"`

	// 目标表名
	TargetTable *string `json:"target_table,omitempty" xml:"target_table"`

	// 字段映射列表
	MappingColumns *[]MappingInfo `json:"mapping_columns,omitempty" xml:"mapping_columns"`
}

func (o MultiTaskMappingCreateBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MultiTaskMappingCreateBody struct{}"
	}

	return strings.Join([]string{"MultiTaskMappingCreateBody", string(data)}, " ")
}
