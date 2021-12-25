package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateEnvironmentVariableV2Request struct {
	// 实例ID

	InstanceId string `json:"instance_id"`
	// 环境变量的编号

	EnvVariableId string `json:"env_variable_id"`

	Body *EnvVariableBase `json:"body,omitempty"`
}

func (o UpdateEnvironmentVariableV2Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateEnvironmentVariableV2Request struct{}"
	}

	return strings.Join([]string{"UpdateEnvironmentVariableV2Request", string(data)}, " ")
}
