package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PollingPolicyRequest struct {

	// 当前轮询的序号。
	Order int32 `json:"order"`

	// 订阅终端urn列表。
	SubscriptionUrns []string `json:"subscription_urns"`
}

func (o PollingPolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PollingPolicyRequest struct{}"
	}

	return strings.Join([]string{"PollingPolicyRequest", string(data)}, " ")
}
