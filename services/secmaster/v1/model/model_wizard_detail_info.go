package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// WizardDetailInfo 页面详情
type WizardDetailInfo struct {

	// 创建时间
	CreateTime *string `json:"create_time,omitempty"`

	// 创建者ID
	CreatorId *string `json:"creator_id,omitempty"`

	// 描述信息
	Description *string `json:"description,omitempty"`

	// 英文描述
	EnDescription *string `json:"en_description,omitempty"`

	// 页面ID
	Id *string `json:"id,omitempty"`

	// 布局页面相关信息
	WizardJson *string `json:"wizard_json,omitempty"`

	// 页面名称
	Name *string `json:"name,omitempty"`

	// 英文名称
	EnName *string `json:"en_name,omitempty"`

	// 更新时间
	UpdateTime *string `json:"update_time,omitempty"`

	// 租户ID
	ProjectId *string `json:"project_id,omitempty"`

	// 工作空间ID
	WorkspaceId *string `json:"workspace_id,omitempty"`

	// 是否已绑定按钮
	IsBinding *bool `json:"is_binding,omitempty"`

	// 绑定按钮
	BindingButton *[]WizardDetailInfoBindingButton `json:"binding_button,omitempty"`

	// 是否为系统页面
	IsBuiltIn *bool `json:"is_built_in,omitempty"`

	// BOA底座版本
	BoaVersion *string `json:"boa_version,omitempty"`

	// 安全云脑版本
	Version *string `json:"version,omitempty"`
}

func (o WizardDetailInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WizardDetailInfo struct{}"
	}

	return strings.Join([]string{"WizardDetailInfo", string(data)}, " ")
}
