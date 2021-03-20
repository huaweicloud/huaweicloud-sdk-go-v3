package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ShowSubscriptionResponse struct {
	// 订阅ID，用于唯一标识一个订阅，在创建订阅时由物联网平台分配获得。

	SubscriptionId *string `json:"subscription_id,omitempty"`

	Subject *Subject `json:"subject,omitempty"`
	// 订阅的回调地址，用于接收对应资源事件的通知消息。

	Callbackurl *string `json:"callbackurl,omitempty"`
	// 物联网平台推送通知消息时使用的协议通道。使用“http”填充，表示该订阅推送协议通道为http(s)协议。

	Channel        *string `json:"channel,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowSubscriptionResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowSubscriptionResponse struct{}"
	}

	return strings.Join([]string{"ShowSubscriptionResponse", string(data)}, " ")
}
