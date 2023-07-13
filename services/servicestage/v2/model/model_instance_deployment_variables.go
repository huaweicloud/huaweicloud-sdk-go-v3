package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// InstanceDeploymentVariables 变量参数值，模板部署时模版参数信息
type InstanceDeploymentVariables struct {

	// 环境变量
	Environment *interface{} `json:"environment,omitempty"`

	// 组件变量
	Components *interface{} `json:"components,omitempty"`
}

func (o InstanceDeploymentVariables) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InstanceDeploymentVariables struct{}"
	}

	return strings.Join([]string{"InstanceDeploymentVariables", string(data)}, " ")
}
