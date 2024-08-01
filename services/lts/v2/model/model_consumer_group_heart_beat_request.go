package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ConsumerGroupHeartBeatRequest Request Object
type ConsumerGroupHeartBeatRequest struct {

	// 日志组ID，获取方式请参见：获取项目ID，获取账号ID，日志组ID、日志流ID。 缺省值：None 最小长度：36 最大长度：36
	GroupId string `json:"group_id"`

	// 日志流ID，获取方式请参见：获取项目ID，获取账号ID，日志组ID、日志流ID 缺省值：None 最小长度：36 最大长度：36
	StreamId string `json:"stream_id"`

	// 消费组名
	ConsumerGroupName string `json:"consumer_group_name"`

	// 消费者名称
	ConsumerName string `json:"consumer_name"`

	Body *[]string `json:"body,omitempty"`
}

func (o ConsumerGroupHeartBeatRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConsumerGroupHeartBeatRequest struct{}"
	}

	return strings.Join([]string{"ConsumerGroupHeartBeatRequest", string(data)}, " ")
}
