package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type WizardCreateRequestPojo struct {

	// 页面ID
	Id *string `json:"id,omitempty"`

	// 页面名称
	Name *string `json:"name,omitempty"`

	// 描述
	Description *string `json:"description,omitempty"`

	// 布局ID
	LayoutId *string `json:"layout_id,omitempty"`

	// 租户ID
	ProjectId *string `json:"project_id,omitempty"`

	// 工作空间ID
	WorkspaceId *string `json:"workspace_id,omitempty"`

	// 布局页面相关信息
	WizardJson *string `json:"wizard_json,omitempty"`

	// 是否已绑定按钮
	IsBinding *bool `json:"is_binding,omitempty"`

	// 绑定按钮
	BindingButton *[]WizardDetailInfoBindingButton `json:"binding_button,omitempty"`

	// BOA底座版本
	BoaVersion *string `json:"boa_version,omitempty"`
}

func (o WizardCreateRequestPojo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WizardCreateRequestPojo struct{}"
	}

	return strings.Join([]string{"WizardCreateRequestPojo", string(data)}, " ")
}
