package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RefreshSubscriptionRequest Request Object
type RefreshSubscriptionRequest struct {

	// 语言。默认en-us。
	XLanguage *string `json:"X-Language,omitempty"`

	// 实例ID
	InstanceId string `json:"instance_id"`

	// 订阅ID
	SubscriptionId string `json:"subscription_id"`
}

func (o RefreshSubscriptionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RefreshSubscriptionRequest struct{}"
	}

	return strings.Join([]string{"RefreshSubscriptionRequest", string(data)}, " ")
}
