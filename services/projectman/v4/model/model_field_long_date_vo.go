package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// FieldLongDateVo 字段参数返回体
type FieldLongDateVo struct {

	// 字段唯一标识。
	Id *string `json:"id,omitempty"`

	// 字段编码。在项目中使用时一般使用code作为字段标识而不是字段ID。
	Code *string `json:"code,omitempty"`

	// 字段显示名称。
	DisplayName *string `json:"display_name,omitempty"`

	// 字段创建人ID。
	CreatedBy *string `json:"created_by,omitempty"`

	// 字段创建时间。时间戳格式，单位毫秒。
	CreatedDate *int64 `json:"created_date,omitempty"`

	// 字段最后修改人ID。
	ModifiedBy *string `json:"modified_by,omitempty"`

	// 字段最后修改时间。时间戳格式，单位毫秒。
	ModifiedDate *int64 `json:"modified_date,omitempty"`

	// 字段类型标识。
	FieldType *string `json:"field_type,omitempty"`

	// 字段类型ID。用于区分不同的字段类型。
	FieldTypeId *string `json:"field_type_id,omitempty"`

	// 字段类型名称。如单选列表、多选列表、多行文本等。
	FieldTypeName *string `json:"field_type_name,omitempty"`

	// 字段定义类型。用于区分系统字段和自定义字段。
	DefinitionType *string `json:"definition_type,omitempty"`

	// 是否显示在云服务类型的迭代看板卡片模式中。
	ShowOnCard *bool `json:"show_on_card,omitempty"`

	// 字段是否为必填项。
	Optional *bool `json:"optional,omitempty"`

	// 字段是否受控。如果工作项已经基线，修改受控字段值时会触发变更评审。
	Controlled *bool `json:"controlled,omitempty"`

	// 字段是否不可变。更新接口无法更新不可变字段。
	Immutable *bool `json:"immutable,omitempty"`

	// 字段排序序号。数值越小越靠前显示。
	No *int32 `json:"no,omitempty"`

	// 字段默认值。创建工作项时自动填充。
	DefaultValue *string `json:"default_value,omitempty"`

	// 字段选项。单选列表类型字段的选项信息，包含选项ID、编码、显示名称等属性。
	Option *[]OptionEntity `json:"option,omitempty"`

	// 字段所有选项。多选列表类型字段的全部选项信息，数组元素包含选项ID、编码、显示名称等属性。
	AllOptions *[]OptionEntity `json:"all_options,omitempty"`

	// 是否存在同名字段。用于检测字段名称冲突。
	HasSameDisplayName *bool `json:"has_same_display_name,omitempty"`
}

func (o FieldLongDateVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FieldLongDateVo struct{}"
	}

	return strings.Join([]string{"FieldLongDateVo", string(data)}, " ")
}
