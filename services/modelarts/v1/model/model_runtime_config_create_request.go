package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RuntimeConfigCreateRequest **参数解释：** 服务运行配置。  **约束限制：** 不涉及。
type RuntimeConfigCreateRequest struct {
	ServiceInvoke *ServiceInvokeCreateRequest `json:"service_invoke"`

	ServiceLimit *ServiceLimit `json:"service_limit"`

	ServiceSecret *ServiceSecret `json:"service_secret,omitempty"`

	ServerTaskLimit *ServerTaskLimit `json:"server_task_limit,omitempty"`
}

func (o RuntimeConfigCreateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RuntimeConfigCreateRequest struct{}"
	}

	return strings.Join([]string{"RuntimeConfigCreateRequest", string(data)}, " ")
}
