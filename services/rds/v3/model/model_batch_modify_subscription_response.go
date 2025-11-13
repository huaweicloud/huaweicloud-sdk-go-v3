package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchModifySubscriptionResponse Response Object
type BatchModifySubscriptionResponse struct {

	// 修改订阅结果。
	Subscriptions  *[]BatchOperateResponseInfo `json:"subscriptions,omitempty"`
	HttpStatusCode int                         `json:"-"`
}

func (o BatchModifySubscriptionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchModifySubscriptionResponse struct{}"
	}

	return strings.Join([]string{"BatchModifySubscriptionResponse", string(data)}, " ")
}
