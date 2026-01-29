package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteBindingRequest Request Object
type DeleteBindingRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	// vhost名称，名称中包含/时，需要将/替换为__F_SLASH__，否则会调用失败。例如：Vhost名称为/test，入参值为__F_SLASH__test。
	Vhost string `json:"vhost"`

	// Exchange名称
	Exchange string `json:"exchange"`

	// **参数解释**： 绑定目标端类型。 **约束限制**： [不涉及。](tag:sbc,cmcc,tm,hk_tm,ax,hk_sbc,dt)[AMQP版本只支持绑定Queue。](tag:hws,hws_hk,hws_eu) **取值范围**： - Exchange：交换机。 - Queue：队列。 **默认取值**： 不涉及。
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
