package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteBindingRequest Request Object
type DeleteBindingRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	// Vhost名称
	Vhost string `json:"vhost"`

	// Exchange名称
	Exchange string `json:"exchange"`

	// 绑定目标端类型，Exchange或Queue。[（AMQP版本只支持Queue绑定类型）](tag:hws,hws_hk)
	DestinationType string `json:"destination_type"`

	// 绑定的目标端名称
	Destination string `json:"destination"`

	// 绑定路由键，经过URL转译后routing_key，可通过调用[查询Exchange绑定列表](ListBindings.xml)或者[查询指定Queue详情](ShowQueueDetails.xml)接口的响应信息获取。
	PropertiesKey string `json:"properties_key"`
}

func (o DeleteBindingRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteBindingRequest struct{}"
	}

	return strings.Join([]string{"DeleteBindingRequest", string(data)}, " ")
}
