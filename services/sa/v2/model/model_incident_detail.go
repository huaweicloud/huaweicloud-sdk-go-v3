package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type IncidentDetail struct {
	DataObject *Incident `json:"data_object,omitempty"`

	// Create time
	CreateTime *string `json:"create_time,omitempty"`

	// Update time
	UpdateTime *string `json:"update_time,omitempty"`

	// Id value
	ProjectId *string `json:"project_id,omitempty"`

	// Id value
	WorkspaceId *string `json:"workspace_id,omitempty"`

	// Id value
	DataclassId *string `json:"dataclass_id,omitempty"`

	// Id value
	LayoutId *string `json:"layout_id,omitempty"`

	// The name, display only
	Name *string `json:"name,omitempty"`

	// The name, display only
	Type *string `json:"type,omitempty"`

	Dataclass *ShowAlertDetailDataclassRef `json:"dataclass,omitempty"`
}

func (o IncidentDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IncidentDetail struct{}"
	}

	return strings.Join([]string{"IncidentDetail", string(data)}, " ")
}
