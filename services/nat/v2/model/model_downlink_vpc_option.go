package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DownlinkVpcOption 私网NAT网关实例所属的VPC实例。
type DownlinkVpcOption struct {

	// 私网NAT网关实例所属的子网的ID。
	VirsubnetId string `json:"virsubnet_id"`

	// 私网NAT网关的ngport_ip_addrss。
	NgportIpAddress *string `json:"ngport_ip_address,omitempty"`
}

func (o DownlinkVpcOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DownlinkVpcOption struct{}"
	}

	return strings.Join([]string{"DownlinkVpcOption", string(data)}, " ")
}
