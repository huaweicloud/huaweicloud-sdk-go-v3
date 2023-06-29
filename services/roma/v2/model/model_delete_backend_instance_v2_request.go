package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteBackendInstanceV2Request Request Object
type DeleteBackendInstanceV2Request struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	// VPC通道的编号
	VpcChannelId string `json:"vpc_channel_id"`

	// 后端实例对象的编号
	MemberId string `json:"member_id"`
}

func (o DeleteBackendInstanceV2Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteBackendInstanceV2Request struct{}"
	}

	return strings.Join([]string{"DeleteBackendInstanceV2Request", string(data)}, " ")
}
