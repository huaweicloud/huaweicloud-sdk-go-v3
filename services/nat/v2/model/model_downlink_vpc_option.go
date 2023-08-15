package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DownlinkVpcOption 私网NAT网关实例所属的VPC实例。
type DownlinkVpcOption struct {

	// 私网NAT网关实例所属的子网的ID。
	VirsubnetId string `json:"virsubnet_id"`
}

func (o DownlinkVpcOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DownlinkVpcOption struct{}"
	}

	return strings.Join([]string{"DownlinkVpcOption", string(data)}, " ")
}
