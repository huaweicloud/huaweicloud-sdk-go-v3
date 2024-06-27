package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAlarmStatisticsUsingRequest Request Object
type ListAlarmStatisticsUsingRequest struct {

	// 服务id
	ServiceId string `json:"service_id"`

	Body *AlarmStatisticsQuery `json:"body,omitempty"`
}

func (o ListAlarmStatisticsUsingRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAlarmStatisticsUsingRequest struct{}"
	}

	return strings.Join([]string{"ListAlarmStatisticsUsingRequest", string(data)}, " ")
}
