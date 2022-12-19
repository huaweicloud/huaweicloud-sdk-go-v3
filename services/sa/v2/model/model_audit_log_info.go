package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Audit Log Info
type AuditLogInfo struct {

	// instance type.
	InstanceType *string `json:"instance_type,omitempty"`

	// Action id.
	ActionId *string `json:"action_id,omitempty"`

	// action name.
	ActionName *string `json:"action_name,omitempty"`

	// Instance id.
	InstanceId *string `json:"instance_id,omitempty"`

	// parent instance id.
	ParentInstanceId *string `json:"parent_instance_id,omitempty"`

	// log level.
	LogLevel *string `json:"log_level,omitempty"`

	// input.
	Input *string `json:"input,omitempty"`

	// output.
	Output *string `json:"output,omitempty"`

	// error_msg.
	ErrorMsg *string `json:"error_msg,omitempty"`

	// start_time.
	StartTime *string `json:"start_time,omitempty"`

	// end_time.
	EndTime *string `json:"end_time,omitempty"`

	// status.
	Status *string `json:"status,omitempty"`

	// trigger type.
	TriggerType *string `json:"trigger_type,omitempty"`
}

func (o AuditLogInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AuditLogInfo struct{}"
	}

	return strings.Join([]string{"AuditLogInfo", string(data)}, " ")
}
