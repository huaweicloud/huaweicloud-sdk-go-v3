package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// FieldResponseBody 字段详情
type FieldResponseBody struct {

	// Id value
	Id *string `json:"id,omitempty"`

	// 订阅包版本
	CloudPackVersion *string `json:"cloud_pack_version,omitempty"`

	// 关联业务id
	BusinessId *string `json:"business_id,omitempty"`

	// 关联业务
	BusinessType *string `json:"business_type,omitempty"`

	// 数据类名称
	DataclassName *string `json:"dataclass_name,omitempty"`

	// 字段业务编码
	BusinessCode *string `json:"business_code,omitempty"`

	// 字段key
	FieldKey *string `json:"field_key,omitempty"`

	// 字段名称
	Name *string `json:"name,omitempty"`

	// 字段描述
	Description *string `json:"description,omitempty"`

	// 默认值
	DefaultValue *string `json:"default_value,omitempty"`

	// 显示类型
	DisplayType *string `json:"display_type,omitempty"`

	// 字段类型，如shorttext,radio,grid等
	FieldType *string `json:"field_type,omitempty"`

	// 附加json
	ExtraJson *string `json:"extra_json,omitempty"`

	// 工具提示
	FieldTooltip *string `json:"field_tooltip,omitempty"`

	// 输入输出类型
	IuType *string `json:"iu_type,omitempty"`

	// 使用业务
	UsedBy *string `json:"used_by,omitempty"`

	// json模式
	JsonSchema *string `json:"json_schema,omitempty"`

	// 是否内置，true内置，false非内置
	IsBuiltIn *bool `json:"is_built_in,omitempty"`

	// 大小写敏感，true敏感，false不敏感
	CaseSensitive *bool `json:"case_sensitive,omitempty"`

	// 只读模式，true只读，false非只读
	ReadOnly *bool `json:"read_only,omitempty"`

	// 是否必填，true必填，false非必填
	Required *bool `json:"required,omitempty"`

	// 可搜索，true可搜索，false非可搜索
	Searchable *bool `json:"searchable,omitempty"`

	// 可见，true可见，false非可见
	Visible *bool `json:"visible,omitempty"`

	// 可维护，true可维护，false非可维护
	Maintainable *bool `json:"maintainable,omitempty"`

	// 可编辑，true可编辑，false非可编辑
	Editable *bool `json:"editable,omitempty"`

	// 可创建，true可创建，false非可创建
	Creatable *bool `json:"creatable,omitempty"`

	// 是否展示在分类映射外的其他地方
	Mapping *bool `json:"mapping,omitempty"`

	// 目标api
	TargetApi *string `json:"target_api,omitempty"`

	// Creator id value
	CreatorId *string `json:"creator_id,omitempty"`

	// Creator name value
	CreatorName *string `json:"creator_name,omitempty"`

	// Modifier id value
	ModifierId *string `json:"modifier_id,omitempty"`

	// Modifier name value
	ModifierName *string `json:"modifier_name,omitempty"`

	// Create time
	CreateTime *string `json:"create_time,omitempty"`

	// Update time
	UpdateTime *string `json:"update_time,omitempty"`
}

func (o FieldResponseBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FieldResponseBody struct{}"
	}

	return strings.Join([]string{"FieldResponseBody", string(data)}, " ")
}
