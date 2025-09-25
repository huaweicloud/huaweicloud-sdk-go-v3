package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type StartVpecpReq struct {

	// 是否开启内网域名。 - true：开启。 - false：不开启。
	EndpointWithDnsName *bool `json:"endpoint_with_dns_name,omitempty"`

	// 创建专业型终端节点。 - true：开启。 - false：不开启。
	ProfessionVpcep *bool `json:"profession_vpcep,omitempty"`

	// 是否开启IPv4/IPv6双栈网络，仅支持在创建专业型终端节点时开启双栈网络，且集群的VPC子网支持IPv6。 - true：开启。 - false：不开启。
	DualstackEnable *bool `json:"dualstack_enable,omitempty"`
}

func (o StartVpecpReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StartVpecpReq struct{}"
	}

	return strings.Join([]string{"StartVpecpReq", string(data)}, " ")
}
