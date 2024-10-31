package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateServiceInstanceConfig 更新Service instance时输入的配置
type UpdateServiceInstanceConfig struct {
	PgsqlInstanceConfig *PgsqlInstanceConfig `json:"pgsql_instance_config,omitempty"`

	ModelInstanceConfig *UpdateModelServiceInstanceConfig `json:"model_instance_config,omitempty"`
}

func (o UpdateServiceInstanceConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateServiceInstanceConfig struct{}"
	}

	return strings.Join([]string{"UpdateServiceInstanceConfig", string(data)}, " ")
}
