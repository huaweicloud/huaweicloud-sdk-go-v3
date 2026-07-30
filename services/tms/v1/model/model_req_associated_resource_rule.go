package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ReqAssociatedResourceRule struct {

	// 规则配置名称
	SettingName string `json:"setting_name"`

	// 标签键列表
	TagKeys *[]string `json:"tag_keys,omitempty"`

	// 存量资源生效状态：enable、disable
	ExistingResourceStatus *string `json:"existing_resource_status,omitempty"`

	// 关系解除后自动删除能力状态：enable、disable
	AutoDeleteStatus *string `json:"auto_delete_status,omitempty"`

	// 规则生效的regionId
	RegionId string `json:"region_id"`
}

func (o ReqAssociatedResourceRule) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ReqAssociatedResourceRule struct{}"
	}

	return strings.Join([]string{"ReqAssociatedResourceRule", string(data)}, " ")
}
