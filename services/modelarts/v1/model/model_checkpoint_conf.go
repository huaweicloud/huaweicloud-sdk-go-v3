package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CheckpointConf 断点配置信息
type CheckpointConf struct {

	// 断点ID
	CheckpointId *string `json:"checkpoint_id,omitempty"`

	// 保存续训任务的步数。 0：关闭不保，-1：自动无限制。
	SaveCheckpointsMax *int32 `json:"save_checkpoints_max,omitempty"`

	// 跳过步数，0表示不跳过。
	SkippedSteps *int32 `json:"skipped_steps,omitempty"`

	// 是否续训任务。  0：非续训,，1:续训。
	RestoreTraining *int32 `json:"restore_training,omitempty"`
}

func (o CheckpointConf) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckpointConf struct{}"
	}

	return strings.Join([]string{"CheckpointConf", string(data)}, " ")
}
