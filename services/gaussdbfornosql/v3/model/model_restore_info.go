package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RestoreInfo 支持按指定备份文件恢复和按指定时间点恢复。 其中按指定时间点恢复仅支持GeminiDB Cassandra和GeminiDB Influx。
type RestoreInfo struct {

	// 备份文件ID。  用于根据指定备份恢复数据到一个新创建的实例的场景，此场景下该字段取值不能为空。
	BackupId *string `json:"backup_id,omitempty"`

	// 数据恢复参考的指定实例的ID。  用于恢复指定实例的指定时间点的数据到一个新创建的实例的场景，此场景下该字段取值不能为空。
	SourceInstanceId *string `json:"source_instance_id,omitempty"`

	// 数据恢复的指定的时间点。  用于恢复指定实例的指定时间点的数据到一个新创建的实例的场景，此场景下该字段取值不能为空。取值为UTC 13位毫秒数，可通过[查询实例可恢复的时间段]接口进行查询。
	RestoreTime *int64 `json:"restore_time,omitempty"`
}

func (o RestoreInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestoreInfo struct{}"
	}

	return strings.Join([]string{"RestoreInfo", string(data)}, " ")
}
