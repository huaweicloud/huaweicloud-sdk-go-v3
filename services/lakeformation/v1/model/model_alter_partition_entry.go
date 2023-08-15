package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AlterPartitionEntry 修改分区实体
type AlterPartitionEntry struct {
	Partition *PartitionInput `json:"partition"`

	// 原分区值数组
	PartitionValues []string `json:"partition_values"`
}

func (o AlterPartitionEntry) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AlterPartitionEntry struct{}"
	}

	return strings.Join([]string{"AlterPartitionEntry", string(data)}, " ")
}
