package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAlarmStatisticsUsingResponse Response Object
type ListAlarmStatisticsUsingResponse struct {
	Body           *[]AlertStatisticsDto `json:"body,omitempty"`
	HttpStatusCode int                   `json:"-"`
}

func (o ListAlarmStatisticsUsingResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAlarmStatisticsUsingResponse struct{}"
	}

	return strings.Join([]string{"ListAlarmStatisticsUsingResponse", string(data)}, " ")
}
