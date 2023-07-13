package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAlarmDataResponse Response Object
type ListAlarmDataResponse struct {

	// 告警列表。
	AlarmDataList *[]AlarmDataVo `json:"alarm_data_list,omitempty"`

	// 消息总数。
	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListAlarmDataResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAlarmDataResponse struct{}"
	}

	return strings.Join([]string{"ListAlarmDataResponse", string(data)}, " ")
}
