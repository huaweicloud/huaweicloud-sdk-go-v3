package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GetAvailableVpcsRequest Request Object
type GetAvailableVpcsRequest struct {

	// 主实例ID
	InstanceId string `json:"instance_id"`

	// 如果传入vpc_id, 则只返回该VPC下的可用子网
	VpcId *string `json:"vpc_id,omitempty"`

	// 如果传入vpc_name, 则只返回该name对应的VPC下的可用子网
	VpcName *string `json:"vpc_name,omitempty"`
}

func (o GetAvailableVpcsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GetAvailableVpcsRequest struct{}"
	}

	return strings.Join([]string{"GetAvailableVpcsRequest", string(data)}, " ")
}
