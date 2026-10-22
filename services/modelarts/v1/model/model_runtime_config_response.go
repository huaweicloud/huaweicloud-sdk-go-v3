package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RuntimeConfigResponse **参数解释：** 服务运行配置。
type RuntimeConfigResponse struct {
	ServiceInvoke *ServiceInvokeResponse `json:"service_invoke"`

	ServiceLimit *ServiceLimitResponse `json:"service_limit"`

	ServiceSecret *ServiceSecretResponse `json:"service_secret,omitempty"`

	ServerTaskLimit *ServerTaskLimit `json:"server_task_limit,omitempty"`

	// **参数解释：** 智能分流开启字段。 **取值范围：** [true, false]。
	SmartRoutingEnable *bool `json:"smart_routing_enable,omitempty"`
}

func (o RuntimeConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RuntimeConfigResponse struct{}"
	}

	return strings.Join([]string{"RuntimeConfigResponse", string(data)}, " ")
}
