package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// OpenGaussBackupStrategyForResponse 自动备份策略。
type OpenGaussBackupStrategyForResponse struct {

	// 备份时间段。自动备份将在该时间段内触发。  取值范围：非空，格式必须为hh:mm-HH:MM且有效，当前时间指UTC时间。  - HH取值必须比hh大1。 - mm和MM取值必须相同，且取值必须为00。 取值示例：  - 08:00-09:00 - 23:00-00:00
	StartTime string `json:"start_time"`

	// 指定已生成备份文件的可保存天数。  取值范围：1～732。  如果请求体中不填写“backup_strategy”字段，则响应体中 “keep_days”默认返回“7”。
	KeepDays int32 `json:"keep_days"`
}

func (o OpenGaussBackupStrategyForResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OpenGaussBackupStrategyForResponse struct{}"
	}

	return strings.Join([]string{"OpenGaussBackupStrategyForResponse", string(data)}, " ")
}
