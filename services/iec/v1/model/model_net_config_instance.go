package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NetConfigInstance 创建边缘实例网络信息。
type NetConfigInstance struct {

	// 边缘网络ID。
	VpcId string `json:"vpc_id"`

	// 待创建边缘实例子网信息。  需要指定vpcid对应VPC下已创建的子网（subnet）的网络ID，UUID格式。
	Subnets []SubnetConfig `json:"subnets"`
}

func (o NetConfigInstance) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NetConfigInstance struct{}"
	}

	return strings.Join([]string{"NetConfigInstance", string(data)}, " ")
}
