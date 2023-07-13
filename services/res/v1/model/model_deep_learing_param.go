package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeepLearingParam 深度学习作业通用参数
type DeepLearingParam struct {
	InitialParameters *Initial `json:"initial_parameters,omitempty"`

	OptimizeParameters *Optimizer `json:"optimize_parameters,omitempty"`

	RegularParameters *Regular `json:"regular_parameters,omitempty"`

	// 最大迭代轮数。
	MaxIterations *int32 `json:"max_iterations,omitempty"`

	// 提前终止训练轮数。
	EarlyStopIterations *int32 `json:"early_stop_iterations,omitempty"`

	// 批量大小。
	BatchSize *int32 `json:"batch_size,omitempty"`

	// 训练数据集切分数量。
	DatasetSplitParts *int32 `json:"dataset_split_parts,omitempty"`

	// 重新训练。
	RestartTrain *bool `json:"restart_train,omitempty"`
}

func (o DeepLearingParam) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeepLearingParam struct{}"
	}

	return strings.Join([]string{"DeepLearingParam", string(data)}, " ")
}
