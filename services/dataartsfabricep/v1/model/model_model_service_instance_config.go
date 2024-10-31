package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ModelServiceInstanceConfig 模型服务实例的配置，在DeployService中使用
type ModelServiceInstanceConfig struct {
	Resource *ResourceDemand `json:"resource"`
}

func (o ModelServiceInstanceConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModelServiceInstanceConfig struct{}"
	}

	return strings.Join([]string{"ModelServiceInstanceConfig", string(data)}, " ")
}
