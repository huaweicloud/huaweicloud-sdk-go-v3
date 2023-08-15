package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateVpcEndpointRequestBody struct {

	// 选定EP的规格，默认为大规格
	Flavor *string `json:"flavor,omitempty"`

	// 制作EP时使用的租户委托名称
	Xrole *string `json:"xrole,omitempty"`

	// 对接EP使用的租户VPCID
	VpcId string `json:"vpc_id"`

	// 对接EP使用的租户子网ID
	SubnetId string `json:"subnet_id"`
}

func (o CreateVpcEndpointRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateVpcEndpointRequestBody struct{}"
	}

	return strings.Join([]string{"CreateVpcEndpointRequestBody", string(data)}, " ")
}
