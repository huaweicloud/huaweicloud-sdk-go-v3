package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PartitionList struct {

	// 总个数
	TotalCount int64 `json:"total_count"`

	// 分区信息列表
	PartitionInfos []Partition `json:"partition_infos"`
}

func (o PartitionList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PartitionList struct{}"
	}

	return strings.Join([]string{"PartitionList", string(data)}, " ")
}
