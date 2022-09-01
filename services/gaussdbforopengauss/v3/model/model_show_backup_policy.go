package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 备份策略信息。
type ShowBackupPolicy struct {

	// 全量备份文件可以保存的天数。
	KeepDays int32 `json:"keep_days" xml:"keep_days"`

	// 全量备份时间段。自动备份将在该时间段内触发。除关闭自动备份策略外，必选。  取值范围：格式必须为hh:mm-HH:MM且有效，当前时间指UTC时间。  - HH取值必须比hh大1。 - mm和MM取值必须相同，且取值必须为00。
	StartTime string `json:"start_time" xml:"start_time"`

	// 全量备份周期配置。自动备份将在每星期指定的天进行。  取值范围：格式为逗号隔开的数字，数字代表星期。
	Period string `json:"period" xml:"period"`

	// 差量备份周期配置。自动差量备份将每隔周期分钟执行。
	DifferentialPriod *string `json:"differential_priod,omitempty" xml:"differential_priod"`
}

func (o ShowBackupPolicy) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowBackupPolicy struct{}"
	}

	return strings.Join([]string{"ShowBackupPolicy", string(data)}, " ")
}
