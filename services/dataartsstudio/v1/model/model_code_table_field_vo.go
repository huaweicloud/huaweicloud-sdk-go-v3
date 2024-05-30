package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CodeTableFieldVo 码表属性信息。
type CodeTableFieldVo struct {

	// 码表字段ID，填写String类型替代Long类型。
	Id *string `json:"id,omitempty"`

	// 所属码表ID，填写String类型替代Long类型。
	CodeTableId *string `json:"code_table_id,omitempty"`

	// 序号。
	Ordinal int32 `json:"ordinal"`

	// 字段名，英文。
	NameEn string `json:"name_en"`

	// 字段名，中文。
	NameCh string `json:"name_ch"`

	// 描述。
	Description *string `json:"description,omitempty"`

	// 字段类型。
	DataType string `json:"data_type"`

	DomainType *DataTypeDomainEnum `json:"domain_type,omitempty"`

	// 数据类型扩展字段。
	DataTypeExtend *string `json:"data_type_extend,omitempty"`

	// 是否唯一。
	IsUniqueKey *bool `json:"is_unique_key,omitempty"`

	// 码表属性值。
	CodeTableFieldValues *[]CodeTableFieldValueVo `json:"code_table_field_values,omitempty"`

	// 码表属性值总数。
	CountFieldValues *int32 `json:"count_field_values,omitempty"`
}

func (o CodeTableFieldVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CodeTableFieldVo struct{}"
	}

	return strings.Join([]string{"CodeTableFieldVo", string(data)}, " ")
}
