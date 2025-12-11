package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AllowVpcEndpointsOption
type AllowVpcEndpointsOption struct {

	// 描述信息。
	Description *string `json:"description,omitempty"`

	// VPC端点，例如：8di3jdu38d7jhfa7df68adyfiadfia6d。
	VpcEndpointId string `json:"vpc_endpoint_id"`
}

func (o AllowVpcEndpointsOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AllowVpcEndpointsOption struct{}"
	}

	return strings.Join([]string{"AllowVpcEndpointsOption", string(data)}, " ")
}
