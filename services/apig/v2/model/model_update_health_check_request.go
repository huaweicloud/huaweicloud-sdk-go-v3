package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateHealthCheckRequest struct {

	// 实例ID，在API网关控制台的“实例信息”中获取。
	InstanceId string `json:"instance_id"`

	// VPC通道的编号
	VpcChannelId string `json:"vpc_channel_id"`

	Body *VpcHealthConfig `json:"body,omitempty"`
}

func (o UpdateHealthCheckRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateHealthCheckRequest struct{}"
	}

	return strings.Join([]string{"UpdateHealthCheckRequest", string(data)}, " ")
}
