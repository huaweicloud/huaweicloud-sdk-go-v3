package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowRestoreStrategyResponseBody **参数解释**：恢复策略配置，包含恢复类型restore_type的取值及各类型对应的参数。
type ShowRestoreStrategyResponseBody struct {
	RestoreType *FlinkJobRestoreType `json:"restore_type"`

	// **参数解释**：Savepoint的ID，唯一标识一个快照。当restore_type为FROM_SAVEPOINT_ID时返回。 **取值范围**：UUIDv7格式，固定36位，由英文小写字母、数字及中划线组成。
	SavepointId *string `json:"savepoint_id,omitempty"`

	// **参数解释**：Checkpoint的存储路径，指向_metadata文件的上一级目录。仅当restore_type为FROM_CHECKPOINT_PATH时返回。 **取值范围**：1~512个字符，需以obs://开头。
	CheckpointPath *string `json:"checkpoint_path,omitempty"`

	// **参数解释**：允许忽略部分算子状态。仅当restore_type为FROM_SAVEPOINT_ID或FROM_CHECKPOINT_PATH时返回。 **取值范围**：   - true：允许忽略部分算子状态。设为true时，如果快照中包含当前作业不再使用的算子状态，将跳过这些状态而非报错。   - false：不允许忽略算子状态。
	AllowNonRestoredState *bool `json:"allow_non_restored_state,omitempty"`

	// **参数解释**：源表的数据消费起始时间戳，用于指定作业启动后从源表的哪个时间点开始消费数据。仅当restore_type为FROM_SPECIFIED_TIMESTAMP时返回，且仅支持Kafka源表。 **取值范围**：0~4102416000000，13位毫秒时间戳。
	JobStartTimeInMs *int64 `json:"job_start_time_in_ms,omitempty"`
}

func (o ShowRestoreStrategyResponseBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRestoreStrategyResponseBody struct{}"
	}

	return strings.Join([]string{"ShowRestoreStrategyResponseBody", string(data)}, " ")
}
