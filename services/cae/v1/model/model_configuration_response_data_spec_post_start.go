package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ConfigurationResponseDataSpecPostStart 生命周期管理组件配置启动后处理。  Configuration.type为\"lifecycle\"时，返回此参数。
type ConfigurationResponseDataSpecPostStart struct {
	Exec *LifeCycleConfigurationExec `json:"exec,omitempty"`
}

func (o ConfigurationResponseDataSpecPostStart) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConfigurationResponseDataSpecPostStart struct{}"
	}

	return strings.Join([]string{"ConfigurationResponseDataSpecPostStart", string(data)}, " ")
}
