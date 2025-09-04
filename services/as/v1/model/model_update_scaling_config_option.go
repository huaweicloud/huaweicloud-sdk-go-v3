package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateScalingConfigOption struct {
	InstanceConfig *UpdateInstanceConfig `json:"instance_config"`
}

func (o UpdateScalingConfigOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateScalingConfigOption struct{}"
	}

	return strings.Join([]string{"UpdateScalingConfigOption", string(data)}, " ")
}
