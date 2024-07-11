package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateBindingBody struct {

	// 要投递的Exchange或Queue名称
	Destination string `json:"destination"`

	// 绑定键值，用于告知Exchange应该将消息投递到哪些Queue或Exchange中
	RoutingKey string `json:"routing_key"`

	// 绑定目标端类型，Exchange或Queue。[（AMQP版本只支持绑定Queue）](tag:hws,hws_hk)
	DestinationType string `json:"destination_type"`
}

func (o CreateBindingBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateBindingBody struct{}"
	}

	return strings.Join([]string{"CreateBindingBody", string(data)}, " ")
}
