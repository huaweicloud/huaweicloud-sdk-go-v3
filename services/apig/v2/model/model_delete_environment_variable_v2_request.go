package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DeleteEnvironmentVariableV2Request struct {
	// 实例编号

	InstanceId string `json:"instance_id"`
	// 环境变量的ID

	EnvVariableId string `json:"env_variable_id"`
}

func (o DeleteEnvironmentVariableV2Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteEnvironmentVariableV2Request struct{}"
	}

	return strings.Join([]string{"DeleteEnvironmentVariableV2Request", string(data)}, " ")
}
