package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 备份策略。
type OpenGaussBackupStrategy struct {
	// 备份时间段。自动备份将在该时间段内触发。  取值范围：非空，格式必须为hh:mm-HH:MM且有效，当前时间指UTC时间。  - HH取值必须比hh大1。 - mm和MM取值必须相同，且取值必须为00。 取值示例：  - 08:00-09:00 - 23:00-00:00

	StartTime string `json:"start_time"`
	// 指定备份文件的可保存天数。  取值范围：1～732。该参数缺省时，默认填写为7天。

	KeepDays *int32 `json:"keep_days,omitempty"`
}

func (o OpenGaussBackupStrategy) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OpenGaussBackupStrategy struct{}"
	}

	return strings.Join([]string{"OpenGaussBackupStrategy", string(data)}, " ")
}
