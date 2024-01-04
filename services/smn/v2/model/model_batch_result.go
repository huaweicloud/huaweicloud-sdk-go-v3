package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchResult struct {

	// 执行结果code
	Code *string `json:"code,omitempty"`

	// 执行结果message
	Message *string `json:"message,omitempty"`

	// 订阅者urn
	SubscriptionUrn *string `json:"subscription_urn,omitempty"`
}

func (o BatchResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchResult struct{}"
	}

	return strings.Join([]string{"BatchResult", string(data)}, " ")
}
