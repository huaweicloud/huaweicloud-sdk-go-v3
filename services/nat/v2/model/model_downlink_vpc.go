package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DownlinkVpc 私网NAT网关实例所属VPC实例。
type DownlinkVpc struct {

	// 私网NAT网关实例所属VPC的ID。
	VpcId string `json:"vpc_id"`

	// 私网NAT网关实例所属子网的ID。
	VirsubnetId string `json:"virsubnet_id"`
}

func (o DownlinkVpc) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DownlinkVpc struct{}"
	}

	return strings.Join([]string{"DownlinkVpc", string(data)}, " ")
}
