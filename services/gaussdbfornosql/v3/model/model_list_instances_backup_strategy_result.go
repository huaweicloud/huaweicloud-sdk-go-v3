package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListInstancesBackupStrategyResult 自动备份策略。
type ListInstancesBackupStrategyResult struct {

	// 备份时间段。自动备份将在该时间段内触发。当前时间指UTC时间。
	StartTime string `json:"start_time"`

	// 已生成备份文件可以保存的天数。取值范围：0~35。
	KeepDays int32 `json:"keep_days"`
}

func (o ListInstancesBackupStrategyResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInstancesBackupStrategyResult struct{}"
	}

	return strings.Join([]string{"ListInstancesBackupStrategyResult", string(data)}, " ")
}
