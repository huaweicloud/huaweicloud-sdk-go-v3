package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateSubscriptionOrderRequest Request Object
type UpdateSubscriptionOrderRequest struct {

	// 用户当前语言环境
	XLanguage string `json:"X-Language"`

	Body *UpdateOrderInfoReq `json:"body,omitempty"`
}

func (o UpdateSubscriptionOrderRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSubscriptionOrderRequest struct{}"
	}

	return strings.Join([]string{"UpdateSubscriptionOrderRequest", string(data)}, " ")
}
