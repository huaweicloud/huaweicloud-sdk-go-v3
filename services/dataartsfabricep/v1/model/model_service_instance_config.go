package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ServiceInstanceConfig 启动Service实例时输入的配置，在DeployService中使用
type ServiceInstanceConfig struct {
	PgsqlInstanceConfig *PgsqlInstanceConfig `json:"pgsql_instance_config,omitempty"`

	ModelInstanceConfig *ModelServiceInstanceConfig `json:"model_instance_config,omitempty"`
}

func (o ServiceInstanceConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ServiceInstanceConfig struct{}"
	}

	return strings.Join([]string{"ServiceInstanceConfig", string(data)}, " ")
}
