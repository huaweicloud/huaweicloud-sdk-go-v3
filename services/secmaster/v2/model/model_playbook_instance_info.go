package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PlaybookInstanceInfo Response of listing playbook instance informations
type PlaybookInstanceInfo struct {

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
}

func (o PlaybookInstanceInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PlaybookInstanceInfo struct{}"
	}

	return strings.Join([]string{"PlaybookInstanceInfo", string(data)}, " ")
}
