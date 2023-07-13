package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Mqtt MQTT配置。 当enable_mqtt取值为false时，mqtts需要为空数组。
type Mqtt struct {

	// MQTT监听地址，根据type取值确定。
	BindAddr string `json:"bind_addr"`

	// 端口号。
	Port int32 `json:"port"`

	// 类型。枚举值： - nic：网卡类型 - ip：IP类型
	Type string `json:"type"`
}

func (o Mqtt) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Mqtt struct{}"
	}

	return strings.Join([]string{"Mqtt", string(data)}, " ")
}
