package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDetailsOfEnvironmentVariableV2Request Request Object
type ShowDetailsOfEnvironmentVariableV2Request struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	// 环境变量的编号
	EnvVariableId string `json:"env_variable_id"`
}

func (o ShowDetailsOfEnvironmentVariableV2Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDetailsOfEnvironmentVariableV2Request struct{}"
	}

	return strings.Join([]string{"ShowDetailsOfEnvironmentVariableV2Request", string(data)}, " ")
}
