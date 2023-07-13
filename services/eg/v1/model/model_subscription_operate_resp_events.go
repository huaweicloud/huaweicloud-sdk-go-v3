package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SubscriptionOperateRespEvents struct {

	// 失败的错误码
	ErrorCode *string `json:"error_code,omitempty"`

	// 失败的原因
	ErrorMsg *string `json:"error_msg,omitempty"`

	// 订阅ID
	SubscriptionId *string `json:"subscription_id,omitempty"`
}

func (o SubscriptionOperateRespEvents) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SubscriptionOperateRespEvents struct{}"
	}

	return strings.Join([]string{"SubscriptionOperateRespEvents", string(data)}, " ")
}
