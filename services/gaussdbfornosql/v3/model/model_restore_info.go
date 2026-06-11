package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RestoreInfo **参数解释：** 备份信息。 **约束限制：** 支持按指定备份恢复和按指定实例的指定时间点恢复。 目前仅GeminiDB Cassandra和GeminiDB Influx集群支持按指定实例的指定时间点恢复。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
type RestoreInfo struct {

	// **参数解释：** 全量备份文件ID。 **约束限制：** 用于根据指定备份恢复数据到一个新创建的实例的场景，此场景下该字段取值不能为空。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	BackupId *string `json:"backup_id,omitempty"`

	// **参数解释：** 数据恢复参考的指定实例的ID。 **约束限制：** 用于恢复指定实例的指定时间点的数据到一个新创建的实例的场景，此场景下该字段取值不能为空。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	SourceInstanceId *string `json:"source_instance_id,omitempty"`

	// **参数解释：** 数据恢复的指定的时间点。 **约束限制：** 用于恢复指定实例的指定时间点的数据到一个新创建的实例的场景，此场景下该字段取值不能为空。 **取值范围：** 取值为UTC 13位毫秒数，可通过查询实例可恢复的时间段 - QueryingtheTimeWindowWhenaBackupCanBeRestored接口进行查询。 **默认取值：** 不涉及。
	RestoreTime *int64 `json:"restore_time,omitempty"`
}

func (o RestoreInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestoreInfo struct{}"
	}

	return strings.Join([]string{"RestoreInfo", string(data)}, " ")
}
