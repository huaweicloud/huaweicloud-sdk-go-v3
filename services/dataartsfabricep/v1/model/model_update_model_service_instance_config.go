package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateModelServiceInstanceConfig 更新模型服务实例的配置
type UpdateModelServiceInstanceConfig struct {
	Resource *BaseDemand `json:"resource"`
}

func (o UpdateModelServiceInstanceConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateModelServiceInstanceConfig struct{}"
	}

	return strings.Join([]string{"UpdateModelServiceInstanceConfig", string(data)}, " ")
}
