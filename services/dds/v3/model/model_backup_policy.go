package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BackupPolicy 备份策略对象，包括备份保留的天数和备份开始时间。
type BackupPolicy struct {

	// 指定已生成的备份文件可以保存的天数。 取值范围：0～732。取0值，表示关闭自动备份策略。
	KeepDays *string `json:"keep_days,omitempty"`

	// 备份时间段。自动备份将在该时间段内触发。开启自动备份策略时，该参数必选；关闭自动备份策略时，不传该参数。 取值范围：格式必须为hh:mm-HH:MM，且有效，当前时间指UTC时间。 - HH取值必须比hh大1。 - mm和MM取值必须相同，且取值必须为00、15、30或45。 取值示例： - 08:15-09:15 - 23:00-00:00
	StartTime *string `json:"start_time,omitempty"`

	// 备份周期配置。自动备份将在每星期指定的天进行。取值范围：格式为半角逗号隔开的数字，数字代表星期。保留天数取值不同，备份周期约束如下： - 0天，不传该参数。 - 1～6天，备份周期全选，取值为：1,2,3,4,5,6,7。 - 7～732天，备份周期至少选择一周中的一天。示例：1,2,3,4。
	Period *string `json:"period,omitempty"`

	// 增量备份开关。不传这个参数则不对增备状态进行改动。开启增量备份后，系统会自动进行增量备份。增量备份开关的取值和约束如下： - false，代表关闭增量备份，关闭会将之前做的增量备份清理。 - true，代表开启增量备份，开启增量备份会触发一次全量备份。
	EnableIncrementalBackup *bool `json:"enable_incremental_backup,omitempty"`
}

func (o BackupPolicy) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BackupPolicy struct{}"
	}

	return strings.Join([]string{"BackupPolicy", string(data)}, " ")
}
