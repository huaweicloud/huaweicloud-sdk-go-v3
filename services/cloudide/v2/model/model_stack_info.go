package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type StackInfo struct {
	ArmConfig *StacksConfig `json:"arm_config,omitempty" xml:"arm_config"`

	// bundleUrl
	BundleUrl *string `json:"bundle_url,omitempty" xml:"bundle_url"`

	Config *StacksConfig `json:"config,omitempty" xml:"config"`

	// 创建时间
	CreatedTime *sdktime.SdkTime `json:"created_time,omitempty" xml:"created_time"`

	// 是否删除
	Delete *bool `json:"delete,omitempty" xml:"delete"`

	// 描述
	Description *string `json:"description,omitempty" xml:"description"`

	// 是否可用
	Disable *bool `json:"disable,omitempty" xml:"disable"`

	// 显示名称
	DisplayName *string `json:"display_name,omitempty" xml:"display_name"`

	// id
	Id *int64 `json:"id,omitempty" xml:"id"`

	// 标签
	Label *string `json:"label,omitempty" xml:"label"`

	// 图标
	Logo *string `json:"logo,omitempty" xml:"logo"`

	// region
	Region *string `json:"region,omitempty" xml:"region"`

	// 是否显示
	Show *bool `json:"show,omitempty" xml:"show"`

	// 技术栈名称
	StackName *string `json:"stack_name,omitempty" xml:"stack_name"`

	// tags
	Tags *[]string `json:"tags,omitempty" xml:"tags"`

	// 修改时间
	UpdatedTime *sdktime.SdkTime `json:"updated_time,omitempty" xml:"updated_time"`

	// 使用者
	Users *[]string `json:"users,omitempty" xml:"users"`
}

func (o StackInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StackInfo struct{}"
	}

	return strings.Join([]string{"StackInfo", string(data)}, " ")
}
