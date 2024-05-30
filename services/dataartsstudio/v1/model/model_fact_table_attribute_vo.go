package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// FactTableAttributeVo 事实表维度信息。
type FactTableAttributeVo struct {

	// 编码，填写String类型替代Long类型。
	Id *string `json:"id,omitempty"`

	// 所属事实表ID，只读，填写String类型替代Long类型。
	FactLogicTableId *string `json:"fact_logic_table_id,omitempty"`

	// 序号。
	Ordinal int32 `json:"ordinal"`

	// 维度ID，填写String类型替代Long类型。
	DimensionId *string `json:"dimension_id,omitempty"`

	// 维度角色。
	Role *string `json:"role,omitempty"`

	Dimension *DimensionVo `json:"dimension,omitempty"`

	// 是否主键。
	IsPrimaryKey bool `json:"is_primary_key"`

	// 是否分区键。
	IsPartitionKey bool `json:"is_partition_key"`

	// 是否外键，只读。
	IsForeignKey *bool `json:"is_foreign_key,omitempty"`

	// 密级
	SecrecyLevels *[]SecrecyLevelRecordVo `json:"secrecy_levels,omitempty"`

	// 描述。
	Description *string `json:"description,omitempty"`

	// 字段类型。
	DataType string `json:"data_type"`

	DomainType *DataTypeDomainEnum `json:"domain_type,omitempty"`

	// 数据类型扩展字段。
	DataTypeExtend *string `json:"data_type_extend,omitempty"`

	// 英文名。
	NameEn string `json:"name_en"`

	// 中文名。
	NameCh string `json:"name_ch"`

	// 是否不为空。
	NotNull *bool `json:"not_null,omitempty"`

	AttributeType *BizTypeEnum `json:"attribute_type,omitempty"`

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
}

func (o FactTableAttributeVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FactTableAttributeVo struct{}"
	}

	return strings.Join([]string{"FactTableAttributeVo", string(data)}, " ")
}
