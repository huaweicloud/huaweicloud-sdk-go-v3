package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Subscription struct {

	// 创建时间。
	CreatedAt *string `json:"created_at,omitempty"`

	// 订阅ID，唯一性标识。创建订阅时，后台自动生成。
	SubscriptionId *string `json:"subscription_id,omitempty"`

	// 订阅的主题。
	TopicUrns []string `json:"topic_urns"`

	// 订阅的主体。
	Entity *string `json:"entity,omitempty"`

	// 订阅的事件。
	Events *[]string `json:"events,omitempty"`
}

func (o Subscription) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Subscription struct{}"
	}

	return strings.Join([]string{"Subscription", string(data)}, " ")
}
