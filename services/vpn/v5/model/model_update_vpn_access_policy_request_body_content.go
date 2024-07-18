package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateVpnAccessPolicyRequestBodyContent struct {

	// 访问策略名称
	Name *string `json:"name,omitempty"`

	// 关联用户组ID
	UserGroupId *string `json:"user_group_id,omitempty"`

	// 访问策略描述
	Description *string `json:"description,omitempty"`

	// 目的IP网段列表，至少有一个网段
	DestIpCidrs *[]string `json:"dest_ip_cidrs,omitempty"`
}

func (o UpdateVpnAccessPolicyRequestBodyContent) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateVpnAccessPolicyRequestBodyContent struct{}"
	}

	return strings.Join([]string{"UpdateVpnAccessPolicyRequestBodyContent", string(data)}, " ")
}
