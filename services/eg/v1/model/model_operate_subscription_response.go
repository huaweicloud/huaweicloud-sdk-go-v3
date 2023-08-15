package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// OperateSubscriptionResponse Response Object
type OperateSubscriptionResponse struct {

	// 操作失败的订阅个数
	FailedCount *int32 `json:"failed_count,omitempty"`

	// 事件列表
	Events *[]SubscriptionOperateRespEvents `json:"events,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o OperateSubscriptionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OperateSubscriptionResponse struct{}"
	}

	return strings.Join([]string{"OperateSubscriptionResponse", string(data)}, " ")
}
