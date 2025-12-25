package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// WizardUpdateRequestPojo 更新页面请求体
type WizardUpdateRequestPojo struct {

	// 创建时间
	CreateTime *string `json:"create_time,omitempty"`

	// 创建者ID
	CreatorId *string `json:"creator_id,omitempty"`

	// 描述信息
	Description *string `json:"description,omitempty"`

	// 页面ID
	Id *string `json:"id,omitempty"`

	// 布局页面相关信息
	WizardJson *string `json:"wizard_json,omitempty"`

	// 页面名称
	Name *string `json:"name,omitempty"`

	// 更新时间
	UpdateTime *string `json:"update_time,omitempty"`

	// 布局ID
	LayoutId *string `json:"layout_id,omitempty"`

	// 租户ID
	ProjectId *string `json:"project_id,omitempty"`

	// 工作空间ID
	WorkspaceId *string `json:"workspace_id,omitempty"`

	// 是否已绑定按钮
	IsBinding *bool `json:"is_binding,omitempty"`

	// 绑定按钮
	BindingButton *[]WizardUpdateRequestPojoBindingButton `json:"binding_button,omitempty"`

	// BOA底座版本
	BoaVersion *string `json:"boa_version,omitempty"`
}

func (o WizardUpdateRequestPojo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WizardUpdateRequestPojo struct{}"
	}

	return strings.Join([]string{"WizardUpdateRequestPojo", string(data)}, " ")
}
