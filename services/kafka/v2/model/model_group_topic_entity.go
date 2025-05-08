package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GroupTopicEntity 消费组TOPIC详情
type GroupTopicEntity struct {

	// TOPIC名称
	Topic *string `json:"topic,omitempty"`

	// 分区
	Partitions *int32 `json:"partitions,omitempty"`

	// 消息堆积数量
	Lag *int32 `json:"lag,omitempty"`
}

func (o GroupTopicEntity) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GroupTopicEntity struct{}"
	}

	return strings.Join([]string{"GroupTopicEntity", string(data)}, " ")
}
