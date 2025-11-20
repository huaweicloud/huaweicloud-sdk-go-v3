package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangeVulScanPolicyRequestInfoTime **参数解释**: 扫描时间
type ChangeVulScanPolicyRequestInfoTime struct {

	// **参数解释**: 扫描日期或者星期列表 **约束限制**: 不涉及 **取值范围**: 最小值0，最大值31 **默认取值**: 不涉及
	DayPart *[]int32 `json:"day_part,omitempty"`

	// **参数解释**: 扫描时间的小时部分 **约束限制**: 不涉及 **取值范围**: 最小值0，最大值11 **默认取值**: 不涉及
	HourPart *int32 `json:"hour_part,omitempty"`

	// **参数解释**: 扫描时间的分钟部分 **约束限制**: 不涉及 **取值范围**: 最小值0，最大值59 **默认取值**: 不涉及
	MinutePart *int32 `json:"minute_part,omitempty"`

	// **参数解释**: 下一次扫描时间 **约束限制**: 不涉及 **取值范围**: 最小值0，最大值2^63-1 **默认取值**: 不涉及
	NextScanTime *int64 `json:"next_scan_time,omitempty"`

	// **参数解释**: 距离标准时区的差值，如东八区（GMT+8）为-480 **约束限制**: 不涉及 **取值范围**: 最小值-840，最大值720 **默认取值**: 不涉及
	TimezoneOffset *int32 `json:"timezone_offset,omitempty"`
}

func (o ChangeVulScanPolicyRequestInfoTime) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeVulScanPolicyRequestInfoTime struct{}"
	}

	return strings.Join([]string{"ChangeVulScanPolicyRequestInfoTime", string(data)}, " ")
}
