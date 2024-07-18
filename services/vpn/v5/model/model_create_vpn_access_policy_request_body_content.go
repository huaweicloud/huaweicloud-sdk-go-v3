package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateVpnAccessPolicyRequestBodyContent struct {

	// 访问策略名称
	Name string `json:"name"`

	// 关联用户组ID
	UserGroupId string `json:"user_group_id"`

	// 访问策略描述
	Description *string `json:"description,omitempty"`

	// 目的IP网段列表，至少有一个网段
	DestIpCidrs []string `json:"dest_ip_cidrs"`
}

func (o CreateVpnAccessPolicyRequestBodyContent) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateVpnAccessPolicyRequestBodyContent struct{}"
	}

	return strings.Join([]string{"CreateVpnAccessPolicyRequestBodyContent", string(data)}, " ")
}
