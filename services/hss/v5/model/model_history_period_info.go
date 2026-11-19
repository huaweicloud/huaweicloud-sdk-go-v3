package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type HistoryPeriodInfo struct {

	// **参数解释**: 报告ID **取值范围**: 最小值0，最大值2147483647
	ReportId *int32 `json:"report_id,omitempty"`

	// **参数解释**: 报告子ID **取值范围**: 最小值0，最大值2147483647
	ReportSubId *int32 `json:"report_sub_id,omitempty"`

	StatPeriod *StatPeriod `json:"stat_period,omitempty"`
}

func (o HistoryPeriodInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HistoryPeriodInfo struct{}"
	}

	return strings.Join([]string{"HistoryPeriodInfo", string(data)}, " ")
}
