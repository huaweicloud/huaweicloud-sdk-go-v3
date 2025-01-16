package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PollingPolicySubscriptionDetails struct {

	// 当前轮询的序号。
	SubscriptionUrn string `json:"subscription_urn"`

	// 订阅终端urn列表。
	Endpoint string `json:"endpoint"`

	// 备注。
	Remark *string `json:"remark,omitempty"`

	// 订阅者状态：0表示订阅还未确认，1表示已经确认，3表示已经取消确认。
	Status *int32 `json:"status,omitempty"`
}

func (o PollingPolicySubscriptionDetails) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PollingPolicySubscriptionDetails struct{}"
	}

	return strings.Join([]string{"PollingPolicySubscriptionDetails", string(data)}, " ")
}
