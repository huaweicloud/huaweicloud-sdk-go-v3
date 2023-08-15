package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAlarmStatisticRequest Request Object
type ListAlarmStatisticRequest struct {

	// 时区
	TimeZone string `json:"time_zone"`
}

func (o ListAlarmStatisticRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAlarmStatisticRequest struct{}"
	}

	return strings.Join([]string{"ListAlarmStatisticRequest", string(data)}, " ")
}
