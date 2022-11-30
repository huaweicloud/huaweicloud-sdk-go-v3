package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowDetailsOfEnvironmentVariableV2Request struct {

	// 实例ID，在API网关控制台的“实例信息”中获取。
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
