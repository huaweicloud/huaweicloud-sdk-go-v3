package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// LayoutFieldResponseBody 字段详情
type LayoutFieldResponseBody struct {

	// 字段ID
	Id *string `json:"id,omitempty"`

	// 订阅包id
	CloudPackId *string `json:"cloud_pack_id,omitempty"`

	// 订阅包名称
	CloudPackName *string `json:"cloud_pack_name,omitempty"`

	// 数据类ID
	DataclassId *string `json:"dataclass_id,omitempty"`

	// 订阅包版本
	CloudPackVersion *string `json:"cloud_pack_version,omitempty"`

	// 字段key
	FieldKey *string `json:"field_key,omitempty"`

	// 字段名称
	Name *string `json:"name,omitempty"`

	// 字段描述
	Description *string `json:"description,omitempty"`

	// 字段英文描述
	EnDescription *string `json:"en_description,omitempty"`

	// 默认值
	DefaultValue *string `json:"default_value,omitempty"`

	// 字段英文默认值
	EnDefaultValue *string `json:"en_default_value,omitempty"`

	// 字段类型，如shorttext,radio,grid等
	FieldType *string `json:"field_type,omitempty"`

	// 附加json
	ExtraJson *string `json:"extra_json,omitempty"`

	// 工具提示
	FieldTooltip *string `json:"field_tooltip,omitempty"`

	// 英文工具提示
	EnFieldTooltip *string `json:"en_field_tooltip,omitempty"`

	// json模式
	JsonSchema *string `json:"json_schema,omitempty"`

	// 是否内置，true内置，false非内置
	IsBuiltIn *bool `json:"is_built_in,omitempty"`

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

	// 创建者ID
	CreatorId *string `json:"creator_id,omitempty"`

	// 创建者名称
	CreatorName *string `json:"creator_name,omitempty"`

	// 修改者ID
	ModifierId *string `json:"modifier_id,omitempty"`

	// 修改者名称
	ModifierName *string `json:"modifier_name,omitempty"`

	// 创建时间
	CreateTime *string `json:"create_time,omitempty"`

	// 更新时间
	UpdateTime *string `json:"update_time,omitempty"`

	// 字段绑定的页面id
	WizardId *string `json:"wizard_id,omitempty"`

	// 字段绑定的流程id
	AopworkflowId *string `json:"aopworkflow_id,omitempty"`

	// 字段绑定的流程版本id
	AopworkflowVersionId *string `json:"aopworkflow_version_id,omitempty"`

	// 字段绑定的剧本id
	PlaybookId *string `json:"playbook_id,omitempty"`

	// 字段绑定的剧本版本id
	PlaybookVersionId *string `json:"playbook_version_id,omitempty"`

	// BOA底座版本
	BoaVersion *string `json:"boa_version,omitempty"`

	// 安全云脑版本
	Version *string `json:"version,omitempty"`
}

func (o LayoutFieldResponseBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LayoutFieldResponseBody struct{}"
	}

	return strings.Join([]string{"LayoutFieldResponseBody", string(data)}, " ")
}
