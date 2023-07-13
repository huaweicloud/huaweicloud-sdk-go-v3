package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DataobjectInfo Information of dataobject
type DataobjectInfo struct {

	// Id value
	Id *string `json:"id,omitempty"`

	// Create time
	CreateTime *string `json:"create_time,omitempty"`

	// Update time
	UpdateTime *string `json:"update_time,omitempty"`

	// Project id value
	ProjectId *string `json:"project_id,omitempty"`

	// dataclass id.
	DataclassId *string `json:"dataclass_id,omitempty"`

	// The name, display only
	Name *string `json:"name,omitempty"`

	// SIMULATION,PLAYBOOK,MANUAL,INSTANCE,DATA_SOURCE
	Type *string `json:"type,omitempty"`

	// data
	Content *string `json:"content,omitempty"`
}

func (o DataobjectInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DataobjectInfo struct{}"
	}

	return strings.Join([]string{"DataobjectInfo", string(data)}, " ")
}
