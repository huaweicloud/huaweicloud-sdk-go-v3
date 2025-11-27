package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RequiredInput struct {

	// 用户名
	Username *string `json:"USERNAME,omitempty"`

	// 控制节点1IP地址
	Master1 *string `json:"MASTER-1,omitempty"`

	// 控制节点2IP地址
	Master2 *string `json:"MASTER-2,omitempty"`

	// 控制节点3IP地址
	Master3 *string `json:"MASTER-3,omitempty"`

	// 是否启用外部负载均衡器
	AccessExternalLoadBalance *bool `json:"ACCESS_EXTERNAL_LOAD_BALANCE,omitempty"`

	// Cilium网络组件的IPv4地址池CIDR
	CiliumIpv4poolCidr *string `json:"CILIUM_IPV4POOL_CIDR,omitempty"`

	// 容器网络的CIDR范围
	NetworkCidr *string `json:"NETWORK_CIDR,omitempty"`

	// DNS服务器IP地址
	DnsServerIp *string `json:"DNS_SERVER_IP,omitempty"`

	// NTP服务器IP地址
	NtpServerIp *string `json:"NTP_SERVER_IP,omitempty"`

	// IAM域ID
	IamDomainId *string `json:"IAM_DOMAIN_ID,omitempty"`
}

func (o RequiredInput) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RequiredInput struct{}"
	}

	return strings.Join([]string{"RequiredInput", string(data)}, " ")
}
