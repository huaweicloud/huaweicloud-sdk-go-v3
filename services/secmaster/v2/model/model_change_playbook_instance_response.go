package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangePlaybookInstanceResponse Response Object
type ChangePlaybookInstanceResponse struct {

	// Id value
	Id *string `json:"id,omitempty"`

	// The name, display only
	Name *string `json:"name,omitempty"`

	// Project id value
	ProjectId *string `json:"project_id,omitempty"`

	Playbook *PlaybookInfoRef `json:"playbook,omitempty"`

	Dataclass *DataclassInfoRef `json:"dataclass,omitempty"`

	Dataobject *DataclassInfoRef `json:"dataobject,omitempty"`

	// Playbook instance status. RUNNING、FINISHED、FAILED、RETRYING、 TERMINATING、TERMINATED
	Status *string `json:"status,omitempty"`

	// trigger type. DEBUG, TIMER, EVENT, MANUAL
	TriggerType *string `json:"trigger_type,omitempty"`

	// Create time
	StartTime *string `json:"start_time,omitempty"`

	// Update time
	EndTime *string `json:"end_time,omitempty"`

	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ChangePlaybookInstanceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangePlaybookInstanceResponse struct{}"
	}

	return strings.Join([]string{"ChangePlaybookInstanceResponse", string(data)}, " ")
}
