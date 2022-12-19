package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Action info
type ActionInfo struct {

	// Id value
	Id *string `json:"id,omitempty"`

	// The name, display only
	Name *string `json:"name,omitempty"`

	// The description, display only
	Description *string `json:"description,omitempty"`

	// Type of action, aopworkflow, Script, Task and so on.
	ActionType *string `json:"action_type,omitempty"`

	// action id
	ActionId *string `json:"action_id,omitempty"`

	// sort_order
	SortOrder *int32 `json:"sort_order,omitempty"`

	// 剧本ID
	PlaybookId *string `json:"playbook_id,omitempty"`

	// 剧本版本ID
	PlaybookVersionId *string `json:"playbook_version_id,omitempty"`

	// project_id
	ProjectId *string `json:"project_id,omitempty"`
}

func (o ActionInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ActionInfo struct{}"
	}

	return strings.Join([]string{"ActionInfo", string(data)}, " ")
}
