package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Source struct {

	// 实例ID，可以调用“查询实例列表和详情”接口获取。如果未申请实例，可以调用“创建实例”接口创建。
	InstanceId string `json:"instance_id"`

	// 表示恢复方式，枚举值： - “backup”，表示使用备份文件恢复，按照此方式恢复时，“type”字段为非必选，“backup_id”必选。 - “timestamp”，表示按时间点恢复，按照此方式恢复时，“type”字段必选，“restore_time”必选。
	Type *string `json:"type,omitempty"`

	// 用于恢复的备份ID。当使用备份文件恢复时需要指定该参数。
	BackupId *string `json:"backup_id,omitempty"`

	// 恢复数据的时间点，格式为UNIX时间戳，单位是毫秒，时区为UTC。
	RestoreTime *string `json:"restore_time,omitempty"`
}

func (o Source) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Source struct{}"
	}

	return strings.Join([]string{"Source", string(data)}, " ")
}
