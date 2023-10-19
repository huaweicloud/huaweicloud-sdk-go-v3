package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ApplicableInstances struct {

	// 实例ID。
	InstanceId *string `json:"instance_id,omitempty"`

	// 实例名称。
	InstanceName *string `json:"instance_name,omitempty"`
}

func (o ApplicableInstances) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApplicableInstances struct{}"
	}

	return strings.Join([]string{"ApplicableInstances", string(data)}, " ")
}
