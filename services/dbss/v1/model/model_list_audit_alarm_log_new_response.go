package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAuditAlarmLogNewResponse Response Object
type ListAuditAlarmLogNewResponse struct {

	// 总条数
	TotalNum *int32 `json:"total_num,omitempty"`

	// 告警列表
	AlarmLog       *[]AlarmLogResponseNewAlarmLog `json:"alarm_log,omitempty"`
	HttpStatusCode int                            `json:"-"`
}

func (o ListAuditAlarmLogNewResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAuditAlarmLogNewResponse struct{}"
	}

	return strings.Join([]string{"ListAuditAlarmLogNewResponse", string(data)}, " ")
}
