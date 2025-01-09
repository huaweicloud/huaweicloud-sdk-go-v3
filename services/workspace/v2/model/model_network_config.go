package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NetworkConfig 网络信息
type NetworkConfig struct {

	// VPC ID
	VpcId *string `json:"vpc_id,omitempty"`

	// VPC名称
	VpcName *string `json:"vpc_name,omitempty"`

	// 业务子网，可以指定返回的网络ID订购桌面
	SubnetIds *[]string `json:"subnet_ids,omitempty"`

	// 后端管理组件占用的子网网段
	ManagementSubnetCidr *string `json:"management_subnet_cidr,omitempty"`

	// subnet_ids所返回的业务子网中,被管理节点所占用的子网id
	ManagementNodeSubnetId *string `json:"management_node_subnet_id,omitempty"`

	// VPC配置信息列表。
	VpcConfigInfos *[]VpcConfigInfo `json:"vpc_config_infos,omitempty"`
}

func (o NetworkConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NetworkConfig struct{}"
	}

	return strings.Join([]string{"NetworkConfig", string(data)}, " ")
}
