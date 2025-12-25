package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DisableAlertRuleResponse Response Object
type DisableAlertRuleResponse struct {

	// UUID
	AlertRuleId *string `json:"alert_rule_id,omitempty"`

	Status *JobStatus `json:"status,omitempty"`

	ProcessStatus  *JobProcessStatus `json:"process_status,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o DisableAlertRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DisableAlertRuleResponse struct{}"
	}

	return strings.Join([]string{"DisableAlertRuleResponse", string(data)}, " ")
}
