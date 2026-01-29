package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteSubscriptionOrderRequest Request Object
type DeleteSubscriptionOrderRequest struct {

	// 用户当前语言环境
	XLanguage string `json:"X-Language"`

	Body *UnsubscribeParam `json:"body,omitempty"`
}

func (o DeleteSubscriptionOrderRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteSubscriptionOrderRequest struct{}"
	}

	return strings.Join([]string{"DeleteSubscriptionOrderRequest", string(data)}, " ")
}
