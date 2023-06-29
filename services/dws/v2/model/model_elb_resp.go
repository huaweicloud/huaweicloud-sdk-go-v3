package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ElbResp 集群ELB的相关信息
type ElbResp struct {

	// 公网ip
	PublicIp *string `json:"public_ip,omitempty"`

	// 内网ip
	PrivateIp *string `json:"private_ip,omitempty"`

	// Elb终端地址
	PrivateEndpoint *string `json:"private_endpoint,omitempty"`

	// Elb名称
	Name *string `json:"name,omitempty"`

	// Elb的ID
	Id *string `json:"id,omitempty"`

	// Elb所属VPC的ID
	VpcId *string `json:"vpc_id,omitempty"`
}

func (o ElbResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ElbResp struct{}"
	}

	return strings.Join([]string{"ElbResp", string(data)}, " ")
}
