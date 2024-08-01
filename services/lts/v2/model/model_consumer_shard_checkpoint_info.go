package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ConsumerShardCheckpointInfo struct {

	// 日志Shard ID
	ShardId string `json:"shard_id"`

	// 游标值
	Checkpoint int64 `json:"checkpoint"`
}

func (o ConsumerShardCheckpointInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConsumerShardCheckpointInfo struct{}"
	}

	return strings.Join([]string{"ConsumerShardCheckpointInfo", string(data)}, " ")
}
