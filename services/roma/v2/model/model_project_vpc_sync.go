package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ProjectVpcSync struct {

	// VPC通道编号
	VpcChannelId *string `json:"vpc_channel_id,omitempty"`

	// 新增关联的实例列表
	InstanceIds *[]string `json:"instance_ids,omitempty"`

	// 是否强制同步，默认不强制同步
	ForceSync *bool `json:"force_sync,omitempty"`
}

func (o ProjectVpcSync) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProjectVpcSync struct{}"
	}

	return strings.Join([]string{"ProjectVpcSync", string(data)}, " ")
}
