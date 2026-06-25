package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SparseBackupPolicy struct {

	// 参数解释：  备份策略ID。  取值范围：  不涉及。
	Id *string `json:"id,omitempty"`

	// 参数解释：  备份周期配置。  取值范围：  格式为“日期 月份 星期”形式的Cron表达式，时区为UTC时区。 日期支持范围为1~31、特殊字符*（表示任意值）、特殊字符L（表示最后一天）。填写1~31或L时支持填写多个，以逗号隔开。 月份支持范围为1~12、特殊字符*（表示任意值）。 星期支持范围为1~7（1表示星期一，2表示星期二，依次类推）、特殊字符*（表示任意值）。填写1~7数字时支持填写多个，以逗号隔开。
	Period *string `json:"period,omitempty"`

	// 参数解释：  备份文件可以保存的天数。  取值范围：  1~3660。
	KeepDays *int32 `json:"keep_days,omitempty"`

	// 参数解释：  全量备份时间段。自动备份将在该时间段内触发。  取值范围：  格式为hh:mm-HH:MM，为UTC时间。 HH的值比hh大1。 mm和MM的值相同，且取值必须为00。
	StartTime *string `json:"start_time,omitempty"`
}

func (o SparseBackupPolicy) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SparseBackupPolicy struct{}"
	}

	return strings.Join([]string{"SparseBackupPolicy", string(data)}, " ")
}
