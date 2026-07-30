package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AssociatedResourceSetting 规则的配置信息
type AssociatedResourceSetting struct {

	// 规则的配置名称
	SettingName *string `json:"setting_name,omitempty"`

	// 主资源
	MasterService *string `json:"master_service,omitempty"`

	// 主资源类型
	MasterResourceType *string `json:"master_resource_type,omitempty"`

	// 关联资源
	AssociatedService *string `json:"associated_service,omitempty"`

	// 关联资源类型
	AssociatedResourceType *string `json:"associated_resource_type,omitempty"`

	// 是否规则是对存量资源生效。
	SupportExistingResource *bool `json:"support_existing_resource,omitempty"`

	// 是否支持关系解除后自动删除标签。
	SupportAutoDelete *bool `json:"support_auto_delete,omitempty"`

	// 规则配置支持的区域Id。
	RegionIds *[]string `json:"region_ids,omitempty"`
}

func (o AssociatedResourceSetting) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssociatedResourceSetting struct{}"
	}

	return strings.Join([]string{"AssociatedResourceSetting", string(data)}, " ")
}
