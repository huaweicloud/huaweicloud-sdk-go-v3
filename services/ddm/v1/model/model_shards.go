package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Shards struct {

	// 物理分片所在RDS的ID。
	DataNodeId *string `json:"data_node_id,omitempty"`

	// 物理分片名称。
	PhysicalDbName *string `json:"physical_db_name,omitempty"`

	// 物理分片运行状态。
	Status *string `json:"status,omitempty"`

	// 物理分片序号。
	ShardIndex *int32 `json:"shard_index,omitempty"`
}

func (o Shards) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Shards struct{}"
	}

	return strings.Join([]string{"Shards", string(data)}, " ")
}
