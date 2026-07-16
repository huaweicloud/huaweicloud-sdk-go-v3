package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RuntimeConfig **参数解释：** 服务运行配置。 **约束限制：** 不涉及。
type RuntimeConfig struct {
	ServiceInvoke *ServiceInvoke `json:"service_invoke"`

	ServiceLimit *ServiceLimit `json:"service_limit"`

	ServiceSecret *ServiceSecret `json:"service_secret,omitempty"`

	ServerTaskLimit *ServerTaskLimit `json:"server_task_limit,omitempty"`
}

func (o RuntimeConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RuntimeConfig struct{}"
	}

	return strings.Join([]string{"RuntimeConfig", string(data)}, " ")
}
