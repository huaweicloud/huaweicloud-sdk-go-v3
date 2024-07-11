package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// MigrationRocketMqTopicConfig RocketMQ元数据迁移，RocketMQ Topic元数据。
type MigrationRocketMqTopicConfig struct {

	// Topic名称。
	TopicName *string `json:"topicName,omitempty"`

	// 是否有序消息。
	Order *bool `json:"order,omitempty"`

	// Topic权限。
	Perm *int32 `json:"perm,omitempty"`

	// 读队列个数。
	ReadQueueNums *int32 `json:"readQueueNums,omitempty"`

	// 写队列个数。
	WriteQueueNums *int32 `json:"writeQueueNums,omitempty"`

	// Topic过滤类型。   - SINGLE_TAG：单标签   - MULTI_TAG：多标签
	TopicFilterType *string `json:"topicFilterType,omitempty"`

	// Topic系统标志位。
	TopicSysFlag *int32 `json:"topicSysFlag,omitempty"`
}

func (o MigrationRocketMqTopicConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MigrationRocketMqTopicConfig struct{}"
	}

	return strings.Join([]string{"MigrationRocketMqTopicConfig", string(data)}, " ")
}
