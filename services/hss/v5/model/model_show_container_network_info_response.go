package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowContainerNetworkInfoResponse Response Object
type ShowContainerNetworkInfoResponse struct {

	// 网络模型
	Mode *string `json:"mode,omitempty"`

	// VPC Id
	Vpc *string `json:"vpc,omitempty"`

	// 网络ID
	Subnet *string `json:"subnet,omitempty"`

	// 安全组
	SecurityGroup *string `json:"security_group,omitempty"`

	// IPv4 服务网段
	Ipv4Cidr *string `json:"ipv4_cidr,omitempty"`

	// 容器网络网段
	Cidrs *string `json:"cidrs,omitempty"`

	// 服务转发模式:    - iptables   - ipvs
	KubeProxyMode *string `json:"kube_proxy_mode,omitempty"`

	// 是否支持配置egress规则: true、false
	IsSupportEgress *bool `json:"is_support_egress,omitempty"`
	HttpStatusCode  int   `json:"-"`
}

func (o ShowContainerNetworkInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowContainerNetworkInfoResponse struct{}"
	}

	return strings.Join([]string{"ShowContainerNetworkInfoResponse", string(data)}, " ")
}
