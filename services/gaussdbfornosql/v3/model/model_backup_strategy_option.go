package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 高级备份策略。
type BackupStrategyOption struct {
	// 备份时间段。自动备份将在该时间段内触发。 取值范围：非空，格式必须为hh:mm-HH:MM且有效，当前时间指UTC时间。   - HH取值必须比hh大1。   - mm和MM取值必须相同，且取值必须为00、30。   - 不传该参数，默认的备份时间段为00:00-01:00。   - 取值示例：23:00-00:00。

	StartTime string `json:"start_time"`
	// 指定已生成的备份文件可以保存的天数。 取值范围：0~35。   - 取0值，表示不设置自动备份策略。   - 不传该参数，默认开启自动备份策略，备份文件默认保存7天。

	KeepDays *string `json:"keep_days,omitempty"`
}

func (o BackupStrategyOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BackupStrategyOption struct{}"
	}

	return strings.Join([]string{"BackupStrategyOption", string(data)}, " ")
}
