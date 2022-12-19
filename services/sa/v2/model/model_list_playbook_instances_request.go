package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListPlaybookInstancesRequest struct {

	// application/json;charset=UTF-8
	ContentType string `json:"content-type"`

	// ID of workspace
	WorkspaceId string `json:"workspace_id"`

	// Playbook instance status. RUNNING、FINISHED、FAILED、RETRYING、 TERMINATING、TERMINATED
	Status *string `json:"status,omitempty"`

	// date type, START END
	DateType *string `json:"date_type,omitempty"`

	// name
	Name *string `json:"name,omitempty"`

	// Playbook name.
	PlaybookName *string `json:"playbook_name,omitempty"`

	// Dataclass name.
	DataclassName *string `json:"dataclass_name,omitempty"`

	// Dataobject name.
	DataobjectName *string `json:"dataobject_name,omitempty"`

	// trigger type. DEBUG, TIMER, EVENT, MANUAL
	TriggerType *string `json:"trigger_type,omitempty"`

	// 起始时间
	FromDate *string `json:"from_date,omitempty"`

	// 结束时间
	ToDate *string `json:"to_date,omitempty"`

	// request limit size
	Limit *int32 `json:"limit,omitempty"`

	// request offset, from 0
	Offset *int32 `json:"offset,omitempty"`
}

func (o ListPlaybookInstancesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPlaybookInstancesRequest struct{}"
	}

	return strings.Join([]string{"ListPlaybookInstancesRequest", string(data)}, " ")
}
