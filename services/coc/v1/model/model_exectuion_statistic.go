package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExectuionStatistic 执行实例状态和批次统计
type ExectuionStatistic struct {

	// 执行实例状态
	InstanceStatus string `json:"instance_status"`

	// 该状态下执行实例个数
	InstanceCount int32 `json:"instance_count"`

	// 该状态下批次index列表
	BatchIndexes []int32 `json:"batch_indexes"`
}

func (o ExectuionStatistic) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExectuionStatistic struct{}"
	}

	return strings.Join([]string{"ExectuionStatistic", string(data)}, " ")
}
