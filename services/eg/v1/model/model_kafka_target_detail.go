package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// KafkaTargetDetail 订阅的kafka事件目标参数列表，该字段序列化后总长度不超过1024字节
type KafkaTargetDetail struct {

	// 主题
	Topic string `json:"topic"`

	KeyTransform *KafkaTargetDetailKeyTransform `json:"keyTransform,omitempty"`
}

func (o KafkaTargetDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "KafkaTargetDetail struct{}"
	}

	return strings.Join([]string{"KafkaTargetDetail", string(data)}, " ")
}
