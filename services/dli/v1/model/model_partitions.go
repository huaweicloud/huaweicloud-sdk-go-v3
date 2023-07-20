package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Partitions
type Partitions struct {

	// 总个数
	TotalCount int64 `json:"total_count"`

	// 分区信息列表
	PartitionInfos []PartitionInfos `json:"partition_infos"`
}

func (o Partitions) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Partitions struct{}"
	}

	return strings.Join([]string{"Partitions", string(data)}, " ")
}
