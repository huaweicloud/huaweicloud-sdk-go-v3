package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AssociatedResourceRule 规则信息
type AssociatedResourceRule struct {

	// 规则的配置名称
	SettingName string `json:"setting_name"`

	// 规则作的用标签范围。如果为空则表示对全部标签生效。
	TagKeys *[]string `json:"tag_keys,omitempty"`

	// 特性开关，规则是否在存量资源生效。
	ExistingResourceStatus *string `json:"existing_resource_status,omitempty"`

	// 特性开关，主资源与子资源关系解除后是否自动删除子资源中与主资源标签键一致的标签。
	AutoDeleteStatus *string `json:"auto_delete_status,omitempty"`

	// 规则状态
	Status *string `json:"status,omitempty"`

	// 规则生效的区域Id
	RegionId string `json:"region_id"`
}

func (o AssociatedResourceRule) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssociatedResourceRule struct{}"
	}

	return strings.Join([]string{"AssociatedResourceRule", string(data)}, " ")
}
