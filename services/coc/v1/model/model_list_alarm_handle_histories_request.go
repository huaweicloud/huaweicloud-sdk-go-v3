package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAlarmHandleHistoriesRequest Request Object
type ListAlarmHandleHistoriesRequest struct {

	// 告警id
	AlarmId string `json:"alarm_id"`

	// 偏移量
	Offset int32 `json:"offset"`

	// 每页限制数量
	Limit int32 `json:"limit"`
}

func (o ListAlarmHandleHistoriesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAlarmHandleHistoriesRequest struct{}"
	}

	return strings.Join([]string{"ListAlarmHandleHistoriesRequest", string(data)}, " ")
}
