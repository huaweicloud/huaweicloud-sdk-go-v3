package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateAction struct {

	// The name, display only
	Name *string `json:"name,omitempty"`

	// The description, display only
	Description *string `json:"description,omitempty"`

	// Type of this action, script or aopworkflow.
	ActionType *string `json:"action_type,omitempty"`

	// action id.
	ActionId *string `json:"action_id,omitempty"`

	// sort_order
	SortOrder *string `json:"sort_order,omitempty"`
}

func (o CreateAction) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAction struct{}"
	}

	return strings.Join([]string{"CreateAction", string(data)}, " ")
}
