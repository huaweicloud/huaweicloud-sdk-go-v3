package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 清空分区数据
type TruncatePartitionInput struct {

	// 分区值
	PartitionValues *[][]string `json:"partition_values,omitempty"`
}

func (o TruncatePartitionInput) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TruncatePartitionInput struct{}"
	}

	return strings.Join([]string{"TruncatePartitionInput", string(data)}, " ")
}
