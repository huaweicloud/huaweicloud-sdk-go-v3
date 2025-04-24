package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExecuteSubscriptionOperationResponse Response Object
type ExecuteSubscriptionOperationResponse struct {

	// 操作失败的订阅个数
	FailedCount *int32 `json:"failed_count,omitempty"`

	// 事件列表
	Events *[]SubscriptionOperateRespEvents `json:"events,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ExecuteSubscriptionOperationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecuteSubscriptionOperationResponse struct{}"
	}

	return strings.Join([]string{"ExecuteSubscriptionOperationResponse", string(data)}, " ")
}
