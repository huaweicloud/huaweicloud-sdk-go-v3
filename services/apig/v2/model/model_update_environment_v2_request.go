package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateEnvironmentV2Request struct {
	// 实例ID

	InstanceId string `json:"instance_id"`
	// 环境的ID

	EnvId string `json:"env_id"`

	Body *EnvCreate `json:"body,omitempty"`
}

func (o UpdateEnvironmentV2Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateEnvironmentV2Request struct{}"
	}

	return strings.Join([]string{"UpdateEnvironmentV2Request", string(data)}, " ")
}
