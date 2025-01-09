package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// VpcConfigInfo VPC配置信息。
type VpcConfigInfo struct {

	// 已使用的子网信息。
	UsedSubnetIds *[]string `json:"used_subnet_ids,omitempty"`

	// VPC ID
	VpcId *string `json:"vpc_id,omitempty"`

	// 子网ID列表
	SubnetIds *[]string `json:"subnet_ids,omitempty"`

	// 是否为默认接入VPC
	DefaultAccessVpc *bool `json:"default_access_vpc,omitempty"`
}

func (o VpcConfigInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VpcConfigInfo struct{}"
	}

	return strings.Join([]string{"VpcConfigInfo", string(data)}, " ")
}
