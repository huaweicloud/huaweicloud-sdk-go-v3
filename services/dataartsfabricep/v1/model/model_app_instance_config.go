package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AppInstanceConfig 启动App实例时输入的配置
type AppInstanceConfig struct {
	PgsqlInstanceConfig *PgsqlInstanceConfig `json:"pgsql_instance_config,omitempty"`

	ModelInstanceConfig *ModelServiceInstanceConfig `json:"model_instance_config,omitempty"`
}

func (o AppInstanceConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AppInstanceConfig struct{}"
	}

	return strings.Join([]string{"AppInstanceConfig", string(data)}, " ")
}
