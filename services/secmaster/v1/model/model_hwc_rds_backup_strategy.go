package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// HwcRdsBackupStrategy 备份策略。
type HwcRdsBackupStrategy struct {

	// 备份时间段。自动备份将在该时间段内触发。 当前时间指UTC时间。
	StartTime *string `json:"start_time,omitempty"`

	// 已生成的备份文件可以保存的天数。 取值范围：0～732。为0时，表示未设置自动备份策略或备份策略已关闭。如果需要延长保留时间请联系客服人员申请，自动备份最长可以保留2562天。
	KeepDays *int32 `json:"keep_days,omitempty"`
}

func (o HwcRdsBackupStrategy) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HwcRdsBackupStrategy struct{}"
	}

	return strings.Join([]string{"HwcRdsBackupStrategy", string(data)}, " ")
}
