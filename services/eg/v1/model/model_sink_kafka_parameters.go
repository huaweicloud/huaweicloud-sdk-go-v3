package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SinkKafkaParameters struct {

	// topic名称
	Topic string `json:"topic"`

	// key的转换规则
	KeyTransform *[]TransForm `json:"keyTransform,omitempty"`

	// 目标连接id
	ConnectionId string `json:"connectionId"`
}

func (o SinkKafkaParameters) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SinkKafkaParameters struct{}"
	}

	return strings.Join([]string{"SinkKafkaParameters", string(data)}, " ")
}
