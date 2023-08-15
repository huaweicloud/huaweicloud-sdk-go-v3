package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PolicyoOdCreate struct {

	// 保留日备个数，该备份不受保留最大备份数限制。取值为0到100。若选择该参数，则timezone 也必选。
	DayBackups *int32 `json:"day_backups,omitempty"`

	// 复制的目标项目ID，仅在跨区域复制时才会使用并且必须指定。
	DestinationProjectId *string `json:"destination_project_id,omitempty"`

	// 复制的目标区域，仅在跨区域复制时才会使用并且必须指定。长度限制：0- 255，只能由字母、数字、“_”、“-”组成
	DestinationRegion *string `json:"destination_region,omitempty"`

	// 跨区域复制时，是否启用加速从而缩减复制的时间，如果不指定，默认不启用加速。
	EnableAcceleration *bool `json:"enable_acceleration,omitempty"`

	// 单个备份对象自动备份的最大备份数。取值为-1或0-99999。-1代表不按备份数清理。若该字段和retention_duration_days字段同时为空，备份会永久保留。
	MaxBackups *int32 `json:"max_backups,omitempty"`

	// 保留月备个数，该备份不受保留最大备份数限制。取值为0到100。若选择该参数，则timezone 也必选。
	MonthBackups *int32 `json:"month_backups,omitempty"`

	// 备份保留时长，单位天。最长支持99999天。-1代表不按时间清理。若该字段和max_backups 参数同时为空，备份会永久保留。
	RetentionDurationDays *int32 `json:"retention_duration_days,omitempty"`

	// 用户所在时区,格式形如UTC+08:00, 若选择年备，月备，周备，日备中任一参数，则该参数不能为空。
	Timezone *string `json:"timezone,omitempty"`

	// 保留周备个数，该备份不受保留最大备份数限制。取值为0到100。若选择该参数，则timezone 也必选。
	WeekBackups *int32 `json:"week_backups,omitempty"`

	// 保留年备个数，该备份不受保留最大备份数限制。取值为0到100。若选择该参数，则timezone 也必选。
	YearBackups *int32 `json:"year_backups,omitempty"`

	// 每间隔多少次执行一次全量备份，当取值为 -1 时，不执行全量备份。  最小值：-1  最大值：100
	FullBackupInterval *int32 `json:"full_backup_interval,omitempty"`
}

func (o PolicyoOdCreate) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PolicyoOdCreate struct{}"
	}

	return strings.Join([]string{"PolicyoOdCreate", string(data)}, " ")
}
