package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DimensionLogicTableAttributeVo struct {

	// 维度表ID，填写String类型替代Long类型。
	Id *string `json:"id,omitempty"`

	// 所属维表ID。
	DimensionLogicTableId *string `json:"dimension_logic_table_id,omitempty"`

	// 序号
	Ordinal int32 `json:"ordinal"`

	// 维度属性ID，填写String类型替代Long类型。
	DimensionAttributeId string `json:"dimension_attribute_id"`

	// 字段名，只读。
	NameEn *string `json:"name_en,omitempty"`

	// 业务属性，只读。
	NameCh *string `json:"name_ch,omitempty"`

	// 描述，只读。
	Description *string `json:"description,omitempty"`

	// 字段类型。
	DataType *string `json:"data_type,omitempty"`

	DomainType *DataTypeDomainEnum `json:"domain_type,omitempty"`

	// 数据类型扩展字段。
	DataTypeExtend *string `json:"data_type_extend,omitempty"`

	// 是否主键，只读。
	IsPrimaryKey *bool `json:"is_primary_key,omitempty"`

	// 是否业务主键。
	IsBizPrimary *bool `json:"is_biz_primary,omitempty"`

	// 是否主键分区，只读。
	IsPartitionKey *bool `json:"is_partition_key,omitempty"`

	// 是否不为空。
	NotNull *bool `json:"not_null,omitempty"`

	// 关联的数据标准的ID，填写String类型替代Long类型。
	StandRowId *string `json:"stand_row_id,omitempty"`

	// 关联的数据标准名称，只读。
	StandRowName *string `json:"stand_row_name,omitempty"`

	// 质量信息，只读。
	QualityInfos *[]QualityInfoVo `json:"quality_infos,omitempty"`

	// 别名。
	Alias *string `json:"alias,omitempty"`

	// 自定义项。
	SelfDefinedFields *[]SelfDefinedFieldVo `json:"self_defined_fields,omitempty"`

	// 密级
	SecrecyLevels *[]SecrecyLevelVo `json:"secrecy_levels,omitempty"`
}

func (o DimensionLogicTableAttributeVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DimensionLogicTableAttributeVo struct{}"
	}

	return strings.Join([]string{"DimensionLogicTableAttributeVo", string(data)}, " ")
}
