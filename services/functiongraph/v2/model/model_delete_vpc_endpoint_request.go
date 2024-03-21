package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteVpcEndpointRequest Request Object
type DeleteVpcEndpointRequest struct {

	// vpc ID
	VpcId string `json:"vpc_id"`

	// 子网编号
	SubnetId string `json:"subnet_id"`

	// 消息体的类型（格式）
	ContentType string `json:"Content-Type"`
}

func (o DeleteVpcEndpointRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteVpcEndpointRequest struct{}"
	}

	return strings.Join([]string{"DeleteVpcEndpointRequest", string(data)}, " ")
}
