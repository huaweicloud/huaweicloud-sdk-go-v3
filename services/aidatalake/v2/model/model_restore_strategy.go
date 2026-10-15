package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RestoreStrategy **参数解释**：恢复策略配置，用于指定Flink作业的状态恢复策略。通过restore_type指定恢复类型，不同类型需配合对应参数使用。 **约束限制**：不涉及。
type RestoreStrategy struct {
	RestoreType *FlinkJobRestoreType `json:"restore_type"`

	// **参数解释**：Savepoint的ID，用于唯一标识一个快照，仅当restore_type为FROM_SAVEPOINT_ID时必填。可通过[查询Flink作业Savepoint列表](ListFlinkSavepoint.xml)或[创建Flink作业Savepoint](CreateFlinkSavepoint.xml)接口获取。 **约束限制**：须为已存在的快照ID，传入不存在的快照ID时接口返回404错误。 **取值范围**：UUIDv7格式，固定36位，由英文小写字母、数字及中划线组成。 **默认取值**：不涉及。
	SavepointId *string `json:"savepoint_id,omitempty"`

	// **参数解释**：Checkpoint的存储路径，需指向_metadata文件的上一级目录，仅当restore_type为FROM_CHECKPOINT_PATH时必填。 **约束限制**：不涉及。 **取值范围**：1~512个字符，需以obs://开头。 **默认取值**：不涉及。
	CheckpointPath *string `json:"checkpoint_path,omitempty"`

	// **参数解释**：允许忽略部分算子状态。仅当JAR类型作业且restore_type为FROM_SAVEPOINT_ID或FROM_CHECKPOINT_PATH时必填。 **约束限制**：不涉及。 **取值范围**：   - true：允许忽略部分算子状态。设为true时，如果快照中包含当前作业不再使用的算子状态，将跳过这些状态而非报错。   - false：不允许忽略算子状态。 **默认取值**：不涉及。
	AllowNonRestoredState *bool `json:"allow_non_restored_state,omitempty"`

	// **参数解释**：源表的数据消费起始时间戳，用于指定作业启动后从源表的哪个时间点开始消费数据。仅当restore_type为FROM_SPECIFIED_TIMESTAMP时必填，且仅支持Kafka源表。 **约束限制**：不涉及。 **取值范围**：0~4102416000000，13位毫秒时间戳。 **默认取值**：不涉及。
	JobStartTimeInMs *int64 `json:"job_start_time_in_ms,omitempty"`
}

func (o RestoreStrategy) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestoreStrategy struct{}"
	}

	return strings.Join([]string{"RestoreStrategy", string(data)}, " ")
}
