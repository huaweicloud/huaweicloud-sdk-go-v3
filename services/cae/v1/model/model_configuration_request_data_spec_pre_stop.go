package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ConfigurationRequestDataSpecPreStop 生命周期管理组件配置停止前处理。  type为\"lifecycle\"时，配置此参数。
type ConfigurationRequestDataSpecPreStop struct {
	Exec *LifeCycleConfigurationExec `json:"exec,omitempty"`
}

func (o ConfigurationRequestDataSpecPreStop) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConfigurationRequestDataSpecPreStop struct{}"
	}

	return strings.Join([]string{"ConfigurationRequestDataSpecPreStop", string(data)}, " ")
}
