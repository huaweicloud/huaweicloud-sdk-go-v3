package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowStreamRequest Request Object
type ShowStreamRequest struct {

	// 需要查询的通道名称。
	StreamName string `json:"stream_name"`

	// 从该分区值开始返回分区列表，返回的分区列表不包括此分区。
	StartPartitionId *string `json:"start_partitionId,omitempty"`

	// 单次请求返回的最大分区数。
	LimitPartitions *int32 `json:"limit_partitions,omitempty"`
}

func (o ShowStreamRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowStreamRequest struct{}"
	}

	return strings.Join([]string{"ShowStreamRequest", string(data)}, " ")
}
