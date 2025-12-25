package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EnableAlertRuleResponse Response Object
type EnableAlertRuleResponse struct {

	// UUID
	AlertRuleId *string `json:"alert_rule_id,omitempty"`

	Status *JobStatus `json:"status,omitempty"`

	ProcessStatus  *JobProcessStatus `json:"process_status,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o EnableAlertRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EnableAlertRuleResponse struct{}"
	}

	return strings.Join([]string{"EnableAlertRuleResponse", string(data)}, " ")
}
