package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TableMappingVo struct {

	// 编码
	Id *int64 `json:"id,omitempty"`

	// 名称
	Name string `json:"name"`

	Description *string `json:"description,omitempty"`

	// 目的表id
	TargetTableId *int64 `json:"target_table_id,omitempty"`

	// 来源表所属模型id
	SrcModelId *int64 `json:"src_model_id,omitempty"`

	// 来源模型名称
	SrcModelName *string `json:"src_model_name,omitempty"`

	// 采集的视图来源，dws视图逆向使用
	ViewText *string `json:"view_text,omitempty"`

	// 目的表名称
	TargetTableName *string `json:"target_table_name,omitempty"`

	// 详情
	Details *[]TableMappingDetailVo `json:"details,omitempty"`

	// 映射的表信息
	SourceTables *[]MappingSourceTableVo `json:"source_tables,omitempty"`

	// 映射的字段信息
	SourceFields *[]MappingSourceFieldVo `json:"source_fields,omitempty"`

	// 创建时间
	CreateTime *sdktime.SdkTime `json:"create_time,omitempty"`

	// 更新时间
	UpdateTime *sdktime.SdkTime `json:"update_time,omitempty"`

	// 创建人
	CreateBy *string `json:"create_by,omitempty"`

	// 更新人
	UpdateBy *string `json:"update_by,omitempty"`
}

func (o TableMappingVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TableMappingVo struct{}"
	}

	return strings.Join([]string{"TableMappingVo", string(data)}, " ")
}
