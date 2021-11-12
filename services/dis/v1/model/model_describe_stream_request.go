package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DescribeStreamRequest struct {
	// 需要查询的通道名称。

	StreamName string `json:"stream_name"`
	// 从该分区值开始返回分区列表，返回的分区列表不包括此分区。

	StartPartitionId *string `json:"start_partitionId,omitempty"`
	// 单次请求返回的最大分区数。  取值范围：1~1000。  默认值：100。

	LimitPartitions *int32 `json:"limit_partitions,omitempty"`
}

func (o DescribeStreamRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DescribeStreamRequest struct{}"
	}

	return strings.Join([]string{"DescribeStreamRequest", string(data)}, " ")
}
