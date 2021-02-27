package model

import (
	"encoding/json"

	"strings"
)

type UpdatePartitionCountRequestBody struct {
	// 需要变更分区数量的通道名称。
	StreamName string `json:"stream_name"`
	// 变更的目标分区数量。  取值为大于0的整数。  设置的值大于当前分区数量表示扩容，小于当前分区数量表示缩容。  注意：  每个通道在一小时内扩容和缩容总次数最多5次，且一小时内扩容或缩容操作有一次成功则最近一小时内不允许再次进行扩容或缩容操作。
	TargetPartitionCount int32 `json:"target_partition_count"`
}

func (o UpdatePartitionCountRequestBody) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdatePartitionCountRequestBody struct{}"
	}

	return strings.Join([]string{"UpdatePartitionCountRequestBody", string(data)}, " ")
}
