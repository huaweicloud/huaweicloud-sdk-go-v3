package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// WorkflowServicePackege service package
type WorkflowServicePackege struct {

	// 资源包的UUID。
	PackageId *string `json:"package_id,omitempty"`

	// 资源包状态。
	Status *string `json:"status,omitempty"`

	// 资源池ID。
	PoolId *string `json:"pool_id,omitempty"`

	// 服务ID。
	ServiceId *string `json:"service_id,omitempty"`

	// Workflow工作流ID。
	WorkflowId *string `json:"workflow_id,omitempty"`

	Order *WorkflowPoolOrder `json:"order,omitempty"`

	// 订阅限制。
	ConsumeLimit *int64 `json:"consume_limit,omitempty"`

	// 当前订阅。
	CurrentConsume *int64 `json:"current_consume,omitempty"`

	// 当前时间。
	CurrentDate *string `json:"current_date,omitempty"`

	// 限制标记。
	LimitEnable *bool `json:"limit_enable,omitempty"`

	// 创建时间。
	CreatedAt *string `json:"created_at,omitempty"`
}

func (o WorkflowServicePackege) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WorkflowServicePackege struct{}"
	}

	return strings.Join([]string{"WorkflowServicePackege", string(data)}, " ")
}
