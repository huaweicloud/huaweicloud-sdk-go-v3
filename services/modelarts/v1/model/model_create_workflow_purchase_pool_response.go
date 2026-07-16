package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateWorkflowPurchasePoolResponse Response Object
type CreateWorkflowPurchasePoolResponse struct {

	// 服务包状态。
	Status *string `json:"status,omitempty"`

	// 资源池ID。
	PoolId *string `json:"pool_id,omitempty"`

	// 在线服务ID。
	ServiceId *string `json:"service_id,omitempty"`

	// Workflow工作流ID。
	WorkflowId *string `json:"workflow_id,omitempty"`

	Order *WorkflowPoolOrder `json:"order,omitempty"`

	// 消费限制。
	ConsumeLimit *int32 `json:"consume_limit,omitempty"`

	// 当前消费。
	CurrentConsume *int32 `json:"current_consume,omitempty"`

	// 当前时间。
	CurrentDate *string `json:"current_date,omitempty"`

	// 限制开关。
	LimitEnable *bool `json:"limit_enable,omitempty"`

	// 创建时间。
	CreatedAt *string `json:"created_at,omitempty"`

	// 订阅包的UUID。创建时不需要填，由后台自动生成。
	PackageId      *string `json:"package_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateWorkflowPurchasePoolResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateWorkflowPurchasePoolResponse struct{}"
	}

	return strings.Join([]string{"CreateWorkflowPurchasePoolResponse", string(data)}, " ")
}
