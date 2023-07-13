package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAlarmStatisticResponse Response Object
type ListAlarmStatisticResponse struct {

	// 告警统计列表
	AlarmStatistics *[]AlarmStatisticResponse `json:"alarm_statistics,omitempty"`
	HttpStatusCode  int                       `json:"-"`
}

func (o ListAlarmStatisticResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAlarmStatisticResponse struct{}"
	}

	return strings.Join([]string{"ListAlarmStatisticResponse", string(data)}, " ")
}
