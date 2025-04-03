package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateVgwRequestBodyContent struct {

	// 网关名称
	Name *string `json:"name,omitempty"`

	// 本端子网
	LocalSubnets *[]string `json:"local_subnets,omitempty"`

	// 使能ipv6的本端子网
	LocalSubnetsV6 *[]string `json:"local_subnets_v6,omitempty"`

	// 有效的EIP的ID，表示绑定新的EIP作为双活VPN网关使用的第一个EIP或主备VPN网关的主EIP。
	EipId1 *string `json:"eip_id_1,omitempty"`

	// 有效的EIP的ID，表示绑定新的EIP作为双活VPN网关使用的第二个EIP或主备VPN网关的备EIP。
	EipId2 *string `json:"eip_id_2,omitempty"`

	PolicyTemplate *UpdateRequestPolicyTemplate `json:"policy_template,omitempty"`
}

func (o UpdateVgwRequestBodyContent) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateVgwRequestBodyContent struct{}"
	}

	return strings.Join([]string{"UpdateVgwRequestBodyContent", string(data)}, " ")
}
