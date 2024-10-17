package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAuditAlarmLogResponse Response Object
type ListAuditAlarmLogResponse struct {

	// 总条数
	TotalNum *int32 `json:"total_num,omitempty"`

	// 告警列表
	AlarmLog       *[]AlarmLogResponseAlarmLog `json:"alarm_log,omitempty"`
	HttpStatusCode int                         `json:"-"`
}

func (o ListAuditAlarmLogResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAuditAlarmLogResponse struct{}"
	}

	return strings.Join([]string{"ListAuditAlarmLogResponse", string(data)}, " ")
}
