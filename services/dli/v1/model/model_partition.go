package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Partition struct {

	// 总个数
	TotalCount int64 `json:"total_count"`

	// 分区信息列表
	PartitionInfos []PartitionInfo `json:"partition_infos"`
}

func (o Partition) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Partition struct{}"
	}

	return strings.Join([]string{"Partition", string(data)}, " ")
}
