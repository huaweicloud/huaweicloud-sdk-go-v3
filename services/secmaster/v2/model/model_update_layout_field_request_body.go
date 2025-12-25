package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateLayoutFieldRequestBody 更新布局字段请求体
type UpdateLayoutFieldRequestBody struct {

	// 字段名称
	Name *string `json:"name,omitempty"`

	// 字段描述
	Description *string `json:"description,omitempty"`

	// 布局ID
	LayoutId *string `json:"layout_id,omitempty"`

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

	// json模式
	JsonSchema *string `json:"json_schema,omitempty"`

	// 只读模式，true只读，false非只读
	Readonly *bool `json:"readonly,omitempty"`

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

	// BOA底座版本
	BoaVersion *string `json:"boa_version,omitempty"`
}

func (o UpdateLayoutFieldRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateLayoutFieldRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateLayoutFieldRequestBody", string(data)}, " ")
}
