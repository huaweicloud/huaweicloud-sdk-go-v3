package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PollingPolicyResponse struct {

	// 当前轮询的序号。
	Order *int32 `json:"order,omitempty"`

	// 订阅终端urn列表。
	Subscriptions *[]PollingPolicySubscriptionDetails `json:"subscriptions,omitempty"`
}

func (o PollingPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PollingPolicyResponse struct{}"
	}

	return strings.Join([]string{"PollingPolicyResponse", string(data)}, " ")
}
