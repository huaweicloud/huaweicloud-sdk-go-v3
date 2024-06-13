package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// FlinkRestoreStrategy Flink 作业的启动策略。
type FlinkRestoreStrategy struct {

	// 启动位点类型。 NONE：无状态启动。 LATEST_SAVEPOINT：最新的作业快照启动。 FROM_SAVEPOINT：从指定快照启动。 LATEST_STATE：最新状态启动。
	RestoreType *string `json:"restore_type,omitempty"`

	// 允许忽略部分算子状态。仅当作业选择有状态恢复时需要设置。
	AllowNonRestoredState *bool `json:"allow_non_restored_state,omitempty"`

	// 无状态启动时间，需输入 13 位时间戳。当选择无状态启动时，可以设置本参数让所有支持 startTime 的源表均从该时刻开始读取数据。
	JobStartTimeInMs *int32 `json:"job_start_time_in_ms,omitempty"`

	// 启动作业快照 ID，启动策略为 FROM_SAVEPOINT 时必填。
	SavepointId *string `json:"savepoint_id,omitempty"`
}

func (o FlinkRestoreStrategy) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FlinkRestoreStrategy struct{}"
	}

	return strings.Join([]string{"FlinkRestoreStrategy", string(data)}, " ")
}
