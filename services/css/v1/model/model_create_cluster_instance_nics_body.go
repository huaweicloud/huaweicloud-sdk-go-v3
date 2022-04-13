package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 子网信息。
type CreateClusterInstanceNicsBody struct {
	// 指定虚拟私有云ID，用于集群网络配置。

	VpcId string `json:"vpcId"`
	// 子网ID(网络ID)，其中一个搜索集群所有实例的子网和安全组必须相同。

	NetId string `json:"netId"`
	// 安全组ID，其中一个搜索集群所有实例的子网和安全组必须相同。

	SecurityGroupId string `json:"securityGroupId"`
}

func (o CreateClusterInstanceNicsBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateClusterInstanceNicsBody struct{}"
	}

	return strings.Join([]string{"CreateClusterInstanceNicsBody", string(data)}, " ")
}
