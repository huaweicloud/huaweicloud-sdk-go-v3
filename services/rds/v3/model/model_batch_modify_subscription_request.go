package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchModifySubscriptionRequest Request Object
type BatchModifySubscriptionRequest struct {

	// 语言。默认en-us。
	XLanguage *string `json:"X-Language,omitempty"`

	// 实例ID。
	InstanceId string `json:"instance_id"`

	Body *ModifySubscriptionsRequestBody `json:"body,omitempty"`
}

func (o BatchModifySubscriptionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchModifySubscriptionRequest struct{}"
	}

	return strings.Join([]string{"BatchModifySubscriptionRequest", string(data)}, " ")
}
