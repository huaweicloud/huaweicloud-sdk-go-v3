package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// MasJobConfig 精调训练作业参数
type MasJobConfig struct {

	// 模型ID
	FtJobUuid *string `json:"ft_job_uuid,omitempty"`

	// 模型训练类型
	FtTrainType *string `json:"ft_train_type,omitempty"`

	// 模型类型
	ModelType *string `json:"model_type,omitempty"`

	// 训练作业输出路径
	TrainOutputPath *string `json:"train_output_path,omitempty"`

	// 训练作业进度
	TrainProcess *float64 `json:"train_process,omitempty"`

	// 断点ID
	CheckpointId *string `json:"checkpoint_id,omitempty"`

	TaskEnv *TaskEnv `json:"task_env,omitempty"`

	CheckpointConfig *CheckpointConf `json:"checkpoint_config,omitempty"`
}

func (o MasJobConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MasJobConfig struct{}"
	}

	return strings.Join([]string{"MasJobConfig", string(data)}, " ")
}
