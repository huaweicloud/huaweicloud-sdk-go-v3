package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 深度学习作业通用参数
type DeepLearingParam struct {
	InitialParameters *Initial `json:"initial_parameters,omitempty" xml:"initial_parameters"`

	OptimizeParameters *Optimizer `json:"optimize_parameters,omitempty" xml:"optimize_parameters"`

	RegularParameters *Regular `json:"regular_parameters,omitempty" xml:"regular_parameters"`

	// 最大迭代轮数。
	MaxIterations *int32 `json:"max_iterations,omitempty" xml:"max_iterations"`

	// 提前终止训练轮数。
	EarlyStopIterations *int32 `json:"early_stop_iterations,omitempty" xml:"early_stop_iterations"`

	// 批量大小。
	BatchSize *int32 `json:"batch_size,omitempty" xml:"batch_size"`

	// 训练数据集切分数量。
	DatasetSplitParts *int32 `json:"dataset_split_parts,omitempty" xml:"dataset_split_parts"`

	// 重新训练。
	RestartTrain *bool `json:"restart_train,omitempty" xml:"restart_train"`
}

func (o DeepLearingParam) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeepLearingParam struct{}"
	}

	return strings.Join([]string{"DeepLearingParam", string(data)}, " ")
}
