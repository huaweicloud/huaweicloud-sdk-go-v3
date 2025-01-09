package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAlarmStatisticsRequest Request Object
type ListAlarmStatisticsRequest struct {
}

func (o ListAlarmStatisticsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAlarmStatisticsRequest struct{}"
	}

	return strings.Join([]string{"ListAlarmStatisticsRequest", string(data)}, " ")
}
