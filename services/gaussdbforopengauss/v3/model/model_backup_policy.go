package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 备份策略信息。
type BackupPolicy struct {

	// 备份文件可以保存的天数。  取值范围：1-732天。
	KeepDays int32 `json:"keep_days"`

	// 备份时间段。自动备份将在该时间段内触发。  取值范围：非空，格式必须为hh:mm-HH:MM且有效，当前时间指UTC时间。  HH取值必须比hh大1，mm和MM取值必须相同，且取值必须为00。  取值示例：  21:00-22:00
	StartTime string `json:"start_time"`

	// 全量备份周期配置。自动全量备份将在每星期指定的天进行。取值范围：格式为逗号隔开的数字，数字代表星期。取值示例：1,2,3,4则表示备份周期配置为星期一、星期二、星期三和星期四。
	Period string `json:"period"`

	// 差异备份间隔时间配置。每次自动差异备份的间隔时间。  取值范围：15、30、60、180、360、720、1440。单位：分钟。  取值示例：30
	DifferentialPeriod string `json:"differential_period"`
}

func (o BackupPolicy) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BackupPolicy struct{}"
	}

	return strings.Join([]string{"BackupPolicy", string(data)}, " ")
}
