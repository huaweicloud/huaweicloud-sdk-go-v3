package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAuditAlarmLogRequest Request Object
type ListAuditAlarmLogRequest struct {

	// 实例ID。可在查询实例列表接口的ID字段获取。
	InstanceId string `json:"instance_id"`

	Body *AlarmLogRequest `json:"body,omitempty"`
}

func (o ListAuditAlarmLogRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAuditAlarmLogRequest struct{}"
	}

	return strings.Join([]string{"ListAuditAlarmLogRequest", string(data)}, " ")
}
