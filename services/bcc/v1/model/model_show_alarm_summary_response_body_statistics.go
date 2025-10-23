package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ShowAlarmSummaryResponseBodyStatistics struct {

	// 告警级别，取值范围:1,2,3,4
	Severity string `json:"severity"`

	// 告警数量
	Counts int32 `json:"counts"`
}

func (o ShowAlarmSummaryResponseBodyStatistics) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAlarmSummaryResponseBodyStatistics struct{}"
	}

	return strings.Join([]string{"ShowAlarmSummaryResponseBodyStatistics", string(data)}, " ")
}
