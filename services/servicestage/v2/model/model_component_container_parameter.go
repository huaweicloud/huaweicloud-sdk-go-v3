package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ComponentContainerParameter struct {

	// 容器名称
	Name *string `json:"name,omitempty"`

	Size *ComponentContainerSize `json:"size,omitempty"`

	// 应用环境变量。
	Env *[]ConfigurationEnvParam `json:"env,omitempty"`

	Storage *[]ComponentStorage `json:"storage,omitempty"`

	Lifecycle *ConfigurationLifecycle `json:"lifecycle,omitempty"`

	Probes *ConfigurationProbes `json:"probes,omitempty"`
}

func (o ComponentContainerParameter) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ComponentContainerParameter struct{}"
	}

	return strings.Join([]string{"ComponentContainerParameter", string(data)}, " ")
}
