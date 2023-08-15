package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type GetPartitionsByValuesInput struct {

	// 获取的分区列表每个分区值列表代表一个分区
	Values [][]string `json:"values"`
}

func (o GetPartitionsByValuesInput) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GetPartitionsByValuesInput struct{}"
	}

	return strings.Join([]string{"GetPartitionsByValuesInput", string(data)}, " ")
}
