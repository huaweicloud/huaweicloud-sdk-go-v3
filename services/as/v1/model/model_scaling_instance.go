package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 伸缩实例。
type ScalingInstance struct {

	// 云服务器名称。
	InstanceName *string `json:"instance_name,omitempty" xml:"instance_name"`

	// 云服务器id。
	InstanceId *string `json:"instance_id,omitempty" xml:"instance_id"`

	// 实例伸缩失败原因。
	FailedReason *string `json:"failed_reason,omitempty" xml:"failed_reason"`

	// 实例伸缩失败详情。
	FailedDetails *string `json:"failed_details,omitempty" xml:"failed_details"`

	// 实例配置信息。
	InstanceConfig *string `json:"instance_config,omitempty" xml:"instance_config"`
}

func (o ScalingInstance) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ScalingInstance struct{}"
	}

	return strings.Join([]string{"ScalingInstance", string(data)}, " ")
}
