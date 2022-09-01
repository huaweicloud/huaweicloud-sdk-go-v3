package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type AddingBackendInstancesV2Request struct {

	// 实例ID
	InstanceId string `json:"instance_id" xml:"instance_id"`

	// VPC通道的编号
	VpcChannelId string `json:"vpc_channel_id" xml:"vpc_channel_id"`

	Body *VpcMemberCreate `json:"body,omitempty" xml:"body"`
}

func (o AddingBackendInstancesV2Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddingBackendInstancesV2Request struct{}"
	}

	return strings.Join([]string{"AddingBackendInstancesV2Request", string(data)}, " ")
}
