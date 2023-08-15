package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// InstanceDeployment 部署实例请求体
type InstanceDeployment struct {

	// 实例id
	InstanceId string `json:"instance_id"`

	Variables *InstanceDeploymentVariables `json:"variables,omitempty"`
}

func (o InstanceDeployment) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InstanceDeployment struct{}"
	}

	return strings.Join([]string{"InstanceDeployment", string(data)}, " ")
}
