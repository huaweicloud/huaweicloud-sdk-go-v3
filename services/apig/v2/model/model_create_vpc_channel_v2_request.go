package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateVpcChannelV2Request struct {

	// 实例ID
	InstanceId string `json:"instance_id" xml:"instance_id"`

	Body *VpcCreate `json:"body,omitempty" xml:"body"`
}

func (o CreateVpcChannelV2Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateVpcChannelV2Request struct{}"
	}

	return strings.Join([]string{"CreateVpcChannelV2Request", string(data)}, " ")
}
