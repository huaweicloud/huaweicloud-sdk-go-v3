package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ConfigurationResponseDataSpecPreStop 生命周期管理组件配置停止前处理。  Configuration.type为\"lifecycle\"时，返回此参数。
type ConfigurationResponseDataSpecPreStop struct {
	Exec *LifeCycleConfigurationExec `json:"exec,omitempty"`
}

func (o ConfigurationResponseDataSpecPreStop) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConfigurationResponseDataSpecPreStop struct{}"
	}

	return strings.Join([]string{"ConfigurationResponseDataSpecPreStop", string(data)}, " ")
}
