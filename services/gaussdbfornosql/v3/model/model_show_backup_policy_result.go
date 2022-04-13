package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 备份策略对象，包括备份保留的天数和备份开始时间。
type ShowBackupPolicyResult struct {
	// 备份文件可以保存的天数。

	KeepDays int32 `json:"keep_days"`
	// 备份时间段。自动备份将在该时间段内触发。

	StartTime *string `json:"start_time,omitempty"`
	// 备份周期配置。自动备份将在每星期指定的天进行。

	Period *string `json:"period,omitempty"`
}

func (o ShowBackupPolicyResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowBackupPolicyResult struct{}"
	}

	return strings.Join([]string{"ShowBackupPolicyResult", string(data)}, " ")
}
