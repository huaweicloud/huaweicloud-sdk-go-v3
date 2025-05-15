package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BackupPolicyItem 备份策略对象，包括备份保留的天数和备份开始时间。
type BackupPolicyItem struct {

	// 备份文件可以保存的天数。
	KeepDays int32 `json:"keep_days"`

	// 备份时间段。自动备份将在该时间段内触发。
	StartTime *string `json:"start_time,omitempty"`

	// 备份周期配置。自动备份将在每星期指定的天进行。
	Period *string `json:"period,omitempty"`

	// 是否开启增量备份。true：表示增量备份策略为开启状态；false：表示增量备份策略为关闭状态。
	EnableIncrementalBackup *bool `json:"enable_incremental_backup,omitempty"`
}

func (o BackupPolicyItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BackupPolicyItem struct{}"
	}

	return strings.Join([]string{"BackupPolicyItem", string(data)}, " ")
}
