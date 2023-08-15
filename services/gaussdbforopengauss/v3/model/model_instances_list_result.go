package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type InstancesListResult struct {

	// 实例ID。
	InstanceId *string `json:"instance_id,omitempty"`

	// 实例名称。
	InstanceName *string `json:"instance_name,omitempty"`
}

func (o InstancesListResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InstancesListResult struct{}"
	}

	return strings.Join([]string{"InstancesListResult", string(data)}, " ")
}
