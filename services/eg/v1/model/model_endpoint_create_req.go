package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type EndpointCreateReq struct {

	// endpoint名称，租户下唯一，由字母、数字、点、下划线和中划线组成，必须以字母或数字开头
	Name string `json:"name"`

	// 访问端点所在的VPC的ID
	VpcId string `json:"vpc_id"`

	// 访问端点所在的子网的ID
	SubnetId string `json:"subnet_id"`

	// 描述
	Description *string `json:"description,omitempty"`
}

func (o EndpointCreateReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EndpointCreateReq struct{}"
	}

	return strings.Join([]string{"EndpointCreateReq", string(data)}, " ")
}
