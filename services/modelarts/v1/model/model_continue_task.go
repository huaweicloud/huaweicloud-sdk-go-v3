package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ContinueTask 续训任务信息
type ContinueTask struct {

	// 断点ID
	CheckpointId *string `json:"checkpoint_id,omitempty"`

	// 续训任务模型ID
	SourceModelId *string `json:"source_model_id,omitempty"`

	// 续训任务模型名称
	SourceModelName *string `json:"source_model_name,omitempty"`

	// 轮数。
	Epoch *int32 `json:"epoch,omitempty"`

	// 步数。
	Steps *int32 `json:"steps,omitempty"`

	// 是否最优
	IsBest *bool `json:"is_best,omitempty"`

	// 跳过步数，0表示不跳过。
	SkippedSteps *int32 `json:"skipped_steps,omitempty"`
}

func (o ContinueTask) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ContinueTask struct{}"
	}

	return strings.Join([]string{"ContinueTask", string(data)}, " ")
}
