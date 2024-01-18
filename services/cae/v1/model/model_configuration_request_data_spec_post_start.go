package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ConfigurationRequestDataSpecPostStart 生命周期管理组件配置启动后处理。  ConfigurationItem.type为\"lifecycle\"时，配置此参数。
type ConfigurationRequestDataSpecPostStart struct {
	Exec *LifeCycleConfigurationExec `json:"exec,omitempty"`
}

func (o ConfigurationRequestDataSpecPostStart) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConfigurationRequestDataSpecPostStart struct{}"
	}

	return strings.Join([]string{"ConfigurationRequestDataSpecPostStart", string(data)}, " ")
}
