package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ConfirmSubscriptionRequest Request Object
type ConfirmSubscriptionRequest struct {

	// Topic的唯一的资源标识，可通过[查询主题列表](smn_api_51004.xml)获取该标识。
	TopicUrn *string `json:"topic_urn,omitempty"`

	// 订阅终端地址。
	Endpoint *string `json:"endpoint,omitempty"`

	// 确认订阅Token信息。
	Token string `json:"token"`
}

func (o ConfirmSubscriptionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConfirmSubscriptionRequest struct{}"
	}

	return strings.Join([]string{"ConfirmSubscriptionRequest", string(data)}, " ")
}
