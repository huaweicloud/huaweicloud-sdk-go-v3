package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteAlertRuleResponse Response Object
type DeleteAlertRuleResponse struct {

	// UUID
	AlertRuleId *string `json:"alert_rule_id,omitempty"`

	// 毫秒时间戳
	DeleteTime *int64 `json:"delete_time,omitempty"`

	ProcessStatus  *JobProcessStatus `json:"process_status,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o DeleteAlertRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAlertRuleResponse struct{}"
	}

	return strings.Join([]string{"DeleteAlertRuleResponse", string(data)}, " ")
}
