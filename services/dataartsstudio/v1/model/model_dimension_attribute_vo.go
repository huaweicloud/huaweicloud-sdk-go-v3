package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DimensionAttributeVo struct {

	// 编码
	Id *int64 `json:"id,omitempty"`

	// 维度id
	DimensionId *int64 `json:"dimension_id,omitempty"`

	// 码表属性id
	CodeTableFieldId *int64 `json:"code_table_field_id,omitempty"`

	// 字段名
	NameEn string `json:"name_en"`

	// 业务属性
	NameCh string `json:"name_ch"`

	Description *string `json:"description,omitempty"`

	// 创建人
	CreateBy *string `json:"create_by,omitempty"`

	// 字段类型
	DataType string `json:"data_type"`

	DomainType *DataTypeDomainEnum `json:"domain_type,omitempty"`

	// 数据类型扩展字段
	DataTypeExtend *string `json:"data_type_extend,omitempty"`

	// 是否主键
	IsPrimaryKey bool `json:"is_primary_key"`

	// 是否业务主键
	IsBizPrimary *bool `json:"is_biz_primary,omitempty"`

	// 是否分区
	IsPartitionKey *bool `json:"is_partition_key,omitempty"`

	// 序号
	Ordinal int32 `json:"ordinal"`

	// 是否不为空
	NotNull *bool `json:"not_null,omitempty"`

	// 关联的数据标准的id
	StandRowId *int64 `json:"stand_row_id,omitempty"`

	StandRowName *string `json:"stand_row_name,omitempty"`

	// 质量信息
	QualityInfos *[]QualityInfoVo `json:"quality_infos,omitempty"`

	Status *BizStatusEnum `json:"status,omitempty"`

	// 创建时间
	CreateTime *sdktime.SdkTime `json:"create_time,omitempty"`

	// 更新时间
	UpdateTime *sdktime.SdkTime `json:"update_time,omitempty"`

	// 别名
	Alias *string `json:"alias,omitempty"`

	// 自定义项
	SelfDefinedFields *[]SelfDefinedFieldVo `json:"self_defined_fields,omitempty"`
}

func (o DimensionAttributeVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DimensionAttributeVo struct{}"
	}

	return strings.Join([]string{"DimensionAttributeVo", string(data)}, " ")
}
