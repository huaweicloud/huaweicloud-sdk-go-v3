package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AggregationLogicTableAttributeVo 汇总表属性。
type AggregationLogicTableAttributeVo struct {

	// 编码，填写String类型替代Long类型。
	Id *string `json:"id,omitempty"`

	// 所属汇总表ID，填写String类型替代Long类型。
	AggregationLogicTableId *string `json:"aggregation_logic_table_id,omitempty"`

	// 序号。
	Ordinal *int32 `json:"ordinal,omitempty"`

	// 字段名。
	NameEn string `json:"name_en"`

	// 业务属性。
	NameCh string `json:"name_ch"`

	AttributeType *BizTypeEnum `json:"attribute_type,omitempty"`

	// 是否主键。
	IsPrimaryKey bool `json:"is_primary_key"`

	// 是否分区键。
	IsPartitionKey bool `json:"is_partition_key"`

	// 密级
	SecrecyLevels *[]SecrecyLevelVo `json:"secrecy_levels,omitempty"`

	// 是否不为空。
	NotNull *bool `json:"not_null,omitempty"`

	// 描述。
	Description *string `json:"description,omitempty"`

	// 字段类型。
	DataType string `json:"data_type"`

	DomainType *DataTypeDomainEnum `json:"domain_type,omitempty"`

	// 数据类型扩展字段。
	DataTypeExtend *string `json:"data_type_extend,omitempty"`

	// 属性关联对象的id
	RefId *string `json:"ref_id,omitempty"`

	// 属性关联对象的中文名
	RefNameCh *string `json:"ref_name_ch,omitempty"`

	// 属性关联对象的英文名
	RefNameEn *string `json:"ref_name_en,omitempty"`

	// 关联的数据标准的ID，填写String类型替代Long类型。
	StandRowId *string `json:"stand_row_id,omitempty"`

	// 关联的数据标准名称，只读。
	StandRowName *string `json:"stand_row_name,omitempty"`

	// 质量信息，只读。
	QualityInfos *[]QualityInfoVo `json:"quality_infos,omitempty"`

	// 别名。
	Alias *string `json:"alias,omitempty"`
}

func (o AggregationLogicTableAttributeVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AggregationLogicTableAttributeVo struct{}"
	}

	return strings.Join([]string{"AggregationLogicTableAttributeVo", string(data)}, " ")
}
