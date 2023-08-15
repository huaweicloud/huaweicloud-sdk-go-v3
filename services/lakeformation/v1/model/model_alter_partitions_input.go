package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AlterPartitionsInput 修改分区输入
type AlterPartitionsInput struct {

	// 批量修改分区对象数组
	PartitionInputs []AlterPartitionEntry `json:"partition_inputs"`
}

func (o AlterPartitionsInput) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AlterPartitionsInput struct{}"
	}

	return strings.Join([]string{"AlterPartitionsInput", string(data)}, " ")
}
