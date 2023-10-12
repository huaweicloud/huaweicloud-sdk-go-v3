package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ConfigurationDataSpecPreStop 生命周期管理组件配置停止前处理。
type ConfigurationDataSpecPreStop struct {
	Exec *LifeCycleConfigurationExec `json:"exec,omitempty"`
}

func (o ConfigurationDataSpecPreStop) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConfigurationDataSpecPreStop struct{}"
	}

	return strings.Join([]string{"ConfigurationDataSpecPreStop", string(data)}, " ")
}