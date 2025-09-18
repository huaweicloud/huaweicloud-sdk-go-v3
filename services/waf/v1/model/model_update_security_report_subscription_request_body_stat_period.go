package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateSecurityReportSubscriptionRequestBodyStatPeriod **参数解释：** 统计周期 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
type UpdateSecurityReportSubscriptionRequestBodyStatPeriod struct {
	BeginTime *int32 `json:"begin_time,omitempty"`

	EndTime *int32 `json:"end_time,omitempty"`
}

func (o UpdateSecurityReportSubscriptionRequestBodyStatPeriod) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSecurityReportSubscriptionRequestBodyStatPeriod struct{}"
	}

	return strings.Join([]string{"UpdateSecurityReportSubscriptionRequestBodyStatPeriod", string(data)}, " ")
}
