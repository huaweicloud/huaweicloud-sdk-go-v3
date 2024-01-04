package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListSubscriptionsItem struct {

	// Topic的唯一的资源标识。
	TopicUrn string `json:"topic_urn"`

	// 不同协议对应不同的endpoint（接受消息的接入点）。目前支持的协议包括：  \"default\": 默认协议。  “email”：邮件传输协议，endpoint为邮箱地址。  “sms”：短信传输协议，endpoint为手机号码。  “functionstage”：FunctionGraph（函数）传输协议，endpoint为一个函数。  “http”、“https”：HTTP/HTTPS传输协议，endpoint为URL。
	Protocol string `json:"protocol"`

	// 订阅者的唯一资源标识。
	SubscriptionUrn string `json:"subscription_urn"`

	// Topic创建者的项目ID。
	Owner string `json:"owner"`

	// 接受消息的接入点。
	Endpoint string `json:"endpoint"`

	// 备注。
	Remark string `json:"remark"`

	// 订阅者状态：0表示订阅还未确认，1表示已经确认，3表示已经取消确认。
	Status int32 `json:"status"`

	FilterPolices *[]SubscriptionsFilterPolicy `json:"filter_polices,omitempty"`
}

func (o ListSubscriptionsItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSubscriptionsItem struct{}"
	}

	return strings.Join([]string{"ListSubscriptionsItem", string(data)}, " ")
}
