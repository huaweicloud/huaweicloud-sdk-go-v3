package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CbrPolicyDetailResourceDetail 资源详情
type CbrPolicyDetailResourceDetail struct {

	// 跨区域复制时，是否启用加速从而缩减复制的时间，如果不指定，默认不启用加速。
	EnableAcceleration *bool `json:"enable_acceleration,omitempty"`

	// 每间隔多少次执行一次全量备份，当取值为 -1 时，不执行全量备份。最小值：-1，最大值：100。
	FullBackupInterval *int32 `json:"full_backup_interval,omitempty"`

	// 复制的目标项目ID，仅在跨区域复制时才会使用并且必须指定。
	DestinationProjectId *string `json:"destination_project_id,omitempty"`

	// 数据保护类型
	OperationType *string `json:"operation_type,omitempty"`

	// 用户所在时区,格式形如UTC+08:00。
	Timezone *string `json:"timezone,omitempty"`

	// 保留周备个数，该备份不受保留最大备份数限制。取值为0到100。
	WeekBackups *int32 `json:"week_backups,omitempty"`

	// pattern
	Pattern *string `json:"pattern,omitempty"`

	// 目标region
	DestinationRegion *string `json:"destination_region,omitempty"`

	// 备份保留时长，单位天。最长支持99999天。-1代表不按时间清理。
	RetentionDurationDays *int32 `json:"retention_duration_days,omitempty"`

	// 是否启用
	Enabled *bool `json:"enabled,omitempty"`

	// 保留日备个数，该备份不受保留最大备份数限制。取值为0到100。
	DayBackups *int32 `json:"day_backups,omitempty"`

	// 管理存储库列表
	AssociateVaults *[]BccPolicyAssociateVault `json:"associate_vaults,omitempty"`

	// 开始时间start_time
	StartTime *string `json:"start_time,omitempty"`

	// 最大备份数量max_backups
	MaxBackups *int32 `json:"max_backups,omitempty"`

	// 保留年备个数，该备份不受保留最大备份数限制。取值为0到100。
	YearBackups *int32 `json:"year_backups,omitempty"`

	// 保留月备个数，该备份不受保留最大备份数限制。取值为0到100。
	MonthBackups *int32 `json:"month_backups,omitempty"`

	// 账号备份的账号urn
	CrossAccountUrn *string `json:"cross_account_urn,omitempty"`
}

func (o CbrPolicyDetailResourceDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CbrPolicyDetailResourceDetail struct{}"
	}

	return strings.Join([]string{"CbrPolicyDetailResourceDetail", string(data)}, " ")
}
