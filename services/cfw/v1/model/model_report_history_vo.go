package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ReportHistoryVo struct {

	// **参数解释**： 报告ID **取值范围**： 不涉及
	ReportId *string `json:"report_id,omitempty"`

	// **参数解释**： 发送时间 **取值范围**： 不涉及
	SendTime *int64 `json:"send_time,omitempty"`

	StatisticPeriod *StatisticPeriod `json:"statistic_period,omitempty"`
}

func (o ReportHistoryVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ReportHistoryVo struct{}"
	}

	return strings.Join([]string{"ReportHistoryVo", string(data)}, " ")
}
