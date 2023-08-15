package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ModifyActionInfo Action info
type ModifyActionInfo struct {

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

func (o ModifyActionInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifyActionInfo struct{}"
	}

	return strings.Join([]string{"ModifyActionInfo", string(data)}, " ")
}
