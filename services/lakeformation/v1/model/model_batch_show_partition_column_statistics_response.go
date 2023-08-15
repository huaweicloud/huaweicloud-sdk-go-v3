package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchShowPartitionColumnStatisticsResponse Response Object
type BatchShowPartitionColumnStatisticsResponse struct {

	// 分区统计信息的数量
	FoundPartitionNumber *int32 `json:"found_partition_number,omitempty"`

	// 分区统计信息的列表
	ColumnStatistics map[string][]ColumnStatisticsObj `json:"column_statistics,omitempty"`
	HttpStatusCode   int                              `json:"-"`
}

func (o BatchShowPartitionColumnStatisticsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchShowPartitionColumnStatisticsResponse struct{}"
	}

	return strings.Join([]string{"BatchShowPartitionColumnStatisticsResponse", string(data)}, " ")
}
