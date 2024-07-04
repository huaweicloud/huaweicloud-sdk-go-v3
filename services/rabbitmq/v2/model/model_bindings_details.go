package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BindingsDetails struct {

	// Exchange名称
	Source *string `json:"source,omitempty"`

	// 绑定目标类型
	DestinationType *string `json:"destination_type,omitempty"`

	// 绑定目标的名称
	Destination *string `json:"destination,omitempty"`

	// 绑定键值
	RoutingKey *string `json:"routing_key,omitempty"`

	// 经过URL转译后routing_key
	PropertiesKey *string `json:"properties_key,omitempty"`
}

func (o BindingsDetails) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BindingsDetails struct{}"
	}

	return strings.Join([]string{"BindingsDetails", string(data)}, " ")
}
