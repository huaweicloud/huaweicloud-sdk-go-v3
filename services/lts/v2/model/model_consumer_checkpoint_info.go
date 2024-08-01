package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ConsumerCheckpointInfo struct {

	// 游标值
	Checkpoint *int64 `json:"checkpoint,omitempty"`

	// 消费组名
	ConsumerGroupName *string `json:"consumer_group_name,omitempty"`

	// 消费者名
	ConsumerName *string `json:"consumer_name,omitempty"`

	// 项目ID
	ProjectId *string `json:"project_id,omitempty"`

	// 日志Shard ID
	ShardId *string `json:"shard_id,omitempty"`

	// 更新时间
	UpdateTime *int64 `json:"update_time,omitempty"`
}

func (o ConsumerCheckpointInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConsumerCheckpointInfo struct{}"
	}

	return strings.Join([]string{"ConsumerCheckpointInfo", string(data)}, " ")
}
