package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 备份策略。
type OpenGaussBackupStrategyForListResponse struct {

	// 备份时间段。自动备份将在该时间段内触发。  当前时间指UTC时间。
	StartTime string `json:"start_time"`

	// 已生成的备份文件可以保存的天数。  取值范围：1～732。
	KeepDays int32 `json:"keep_days"`
}

func (o OpenGaussBackupStrategyForListResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OpenGaussBackupStrategyForListResponse struct{}"
	}

	return strings.Join([]string{"OpenGaussBackupStrategyForListResponse", string(data)}, " ")
}
