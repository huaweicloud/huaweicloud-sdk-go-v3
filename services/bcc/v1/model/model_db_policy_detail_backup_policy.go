package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DbPolicyDetailBackupPolicy 数据库备份策略
type DbPolicyDetailBackupPolicy struct {

	// 备份保留时长
	KeepDays *int32 `json:"keep_days,omitempty"`

	// 备份开始时间
	StartTime *string `json:"start_time,omitempty"`

	// 备份周期，一周哪几天
	Period *string `json:"period,omitempty"`

	// gaussdb差异备份周期
	DifferentialPeriod *int32 `json:"differential_period,omitempty"`

	// gaussdb备份限速
	RateLimit *int32 `json:"rate_limit,omitempty"`

	// gaussdb预取页面个数
	PrefetchBlock *int32 `json:"prefetch_block,omitempty"`

	// gaussdb文件拆分大小
	FileSplitSize *int32 `json:"file_split_size,omitempty"`

	// gaussdb是否启用standby备份
	EnableStandbyBackup *bool `json:"enable_standby_backup,omitempty"`
}

func (o DbPolicyDetailBackupPolicy) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DbPolicyDetailBackupPolicy struct{}"
	}

	return strings.Join([]string{"DbPolicyDetailBackupPolicy", string(data)}, " ")
}
