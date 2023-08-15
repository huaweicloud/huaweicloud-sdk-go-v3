package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type InstanceInfoResult struct {

	// 实例ID。
	InstanceId *string `json:"instance_id,omitempty"`

	// 实例名称。
	InstanceName *string `json:"instance_name,omitempty"`

	// 实例状态。
	InstanceStatus *string `json:"instance_status,omitempty"`
}

func (o InstanceInfoResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InstanceInfoResult struct{}"
	}

	return strings.Join([]string{"InstanceInfoResult", string(data)}, " ")
}
