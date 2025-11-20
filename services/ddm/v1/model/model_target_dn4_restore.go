package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TargetDn4Restore struct {

	// 实例名称。
	InstanceName *string `json:"instance_name,omitempty"`

	// 实例ID。
	InstanceId *string `json:"instance_id,omitempty"`

	// 引擎名称。
	EngineName *string `json:"engine_name,omitempty"`

	// 状态。
	Status *string `json:"status,omitempty"`

	// 是否可用。
	Available *bool `json:"available,omitempty"`

	// 无法使用原因。
	UnavailableReason *string `json:"unavailable_reason,omitempty"`

	// 虚拟私有云名称。
	VpcName *string `json:"vpc_name,omitempty"`

	// 私有ip。
	PrivateIp *string `json:"private_ip,omitempty"`
}

func (o TargetDn4Restore) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TargetDn4Restore struct{}"
	}

	return strings.Join([]string{"TargetDn4Restore", string(data)}, " ")
}
