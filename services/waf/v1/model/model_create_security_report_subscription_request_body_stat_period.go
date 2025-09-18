package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateSecurityReportSubscriptionRequestBodyStatPeriod **参数解释：** 统计周期，定义数据统计的时间范围区间 **约束限制：** begin_time 必须小于 end_time，且均为毫秒级时间戳 **默认取值：** 不涉及
type CreateSecurityReportSubscriptionRequestBodyStatPeriod struct {

	// **参数解释：** 开始时间，统计周期的起始时间点（毫秒级时间戳） **格式要求：** 符合Unix时间戳规范，精确到毫秒 **取值范围：** 1970-01-01 00:00:00 UTC至今的时间戳
	BeginTime *int64 `json:"begin_time,omitempty"`

	// **参数解释：** 结束时间，统计周期的终止时间点（毫秒级时间戳） **格式要求：** 符合Unix时间戳规范，精确到毫秒 **取值范围：** 大于begin_time且在1970-01-01 00:00:00 UTC至今的时间戳
	EndTime *int64 `json:"end_time,omitempty"`
}

func (o CreateSecurityReportSubscriptionRequestBodyStatPeriod) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSecurityReportSubscriptionRequestBodyStatPeriod struct{}"
	}

	return strings.Join([]string{"CreateSecurityReportSubscriptionRequestBodyStatPeriod", string(data)}, " ")
}
