package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ContinueTrainTask struct {

	// 中间产物id。
	CheckpointId *string `json:"checkpoint_id,omitempty"`

	// 续训任务id。
	ContinueTaskId *string `json:"continue_task_id,omitempty"`

	// 续训任务名称。
	ContinueTaskName *string `json:"continue_task_name,omitempty"`

	// 续训训练类型。
	ContinueTrainType *string `json:"continue_train_type,omitempty"`

	// 跳过步数，0表示不跳过。
	SkippedSteps *int32 `json:"skipped_steps,omitempty"`

	// 是否续训任务。  0: 非续训, 1:续训。
	RestoreTraining *int32 `json:"restore_training,omitempty"`

	// 创建时间。
	CreateTime *string `json:"create_time,omitempty"`

	// 中间产物配置信息。
	CheckpointConfig *string `json:"checkpoint_config,omitempty"`
}

func (o ContinueTrainTask) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ContinueTrainTask struct{}"
	}

	return strings.Join([]string{"ContinueTrainTask", string(data)}, " ")
}
