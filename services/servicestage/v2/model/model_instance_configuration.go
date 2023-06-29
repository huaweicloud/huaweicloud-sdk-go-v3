package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// InstanceConfiguration 应用配置，环境变量等，如{“env”: [{“name”: “log-level”: “warn”}]}, 默认空。
type InstanceConfiguration struct {

	// 应用环境变量。
	Env *[]ConfigurationEnvParam `json:"env,omitempty"`

	Storage *[]ComponentStorage `json:"storage,omitempty"`

	Strategy *ConfigurationStrategy `json:"strategy,omitempty"`

	Lifecycle *ConfigurationLifecycle `json:"lifecycle,omitempty"`

	Scheduler *ConfigurationScheduler `json:"scheduler,omitempty"`

	Probes *ConfigurationProbes `json:"probes,omitempty"`
}

func (o InstanceConfiguration) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InstanceConfiguration struct{}"
	}

	return strings.Join([]string{"InstanceConfiguration", string(data)}, " ")
}
