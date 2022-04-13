package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/sdktime"
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type StackInfo struct {
	ArmConfig *StacksConfig `json:"arm_config,omitempty"`
	// bundleUrl

	BundleUrl *string `json:"bundle_url,omitempty"`

	Config *StacksConfig `json:"config,omitempty"`
	// 创建时间

	CreatedTime *sdktime.SdkTime `json:"created_time,omitempty"`
	// 是否删除

	Delete *bool `json:"delete,omitempty"`
	// 描述

	Description *string `json:"description,omitempty"`
	// 是否可用

	Disable *bool `json:"disable,omitempty"`
	// 显示名称

	DisplayName *string `json:"display_name,omitempty"`
	// id

	Id *int64 `json:"id,omitempty"`
	// 标签

	Label *string `json:"label,omitempty"`
	// 图标

	Logo *string `json:"logo,omitempty"`
	// region

	Region *string `json:"region,omitempty"`
	// 是否显示

	Show *bool `json:"show,omitempty"`
	// 技术栈名称

	StackName *string `json:"stack_name,omitempty"`
	// tags

	Tags *[]string `json:"tags,omitempty"`
	// 修改时间

	UpdatedTime *sdktime.SdkTime `json:"updated_time,omitempty"`
	// 使用者

	Users *[]string `json:"users,omitempty"`
}

func (o StackInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StackInfo struct{}"
	}

	return strings.Join([]string{"StackInfo", string(data)}, " ")
}
