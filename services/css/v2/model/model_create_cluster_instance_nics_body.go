package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateClusterInstanceNicsBody 子网信息。
type CreateClusterInstanceNicsBody struct {

	// 指定虚拟私有云ID，用于集群网络配置。
	VpcId string `json:"vpcId"`

	// 子网ID(网络ID)。
	NetId string `json:"netId"`

	// 安全组ID。
	SecurityGroupId string `json:"securityGroupId"`

	// 节点IP信息，在指定IP创建集群时配置。
	Ips *[]string `json:"ips,omitempty"`
}

func (o CreateClusterInstanceNicsBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateClusterInstanceNicsBody struct{}"
	}

	return strings.Join([]string{"CreateClusterInstanceNicsBody", string(data)}, " ")
}
