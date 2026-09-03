package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// WorkloadStatisticsInfo 工作负载统计
type WorkloadStatisticsInfo struct {

	// 旧版模型部署作业数量。
	Infer *int32 `json:"infer,omitempty"`

	// 开发环境作业数量。
	Notebook *int32 `json:"notebook,omitempty"`

	// 训练作业数量。
	Train *int32 `json:"train,omitempty"`

	// 权重预热作业数量。
	WarmUpTask *int32 `json:"warmUpTask,omitempty"`

	// 模型部署作业数量。
	XInfer *int32 `json:"x-infer,omitempty"`

	// 所有作业总和。
	Sum *int32 `json:"sum,omitempty"`
}

func (o WorkloadStatisticsInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WorkloadStatisticsInfo struct{}"
	}

	return strings.Join([]string{"WorkloadStatisticsInfo", string(data)}, " ")
}
