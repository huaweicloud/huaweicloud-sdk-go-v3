package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ModifyPlaybookVersionInfo Information of playbook version
type ModifyPlaybookVersionInfo struct {

	// The description, display only
	Description *string `json:"description,omitempty"`

	// 工作空间id
	WorkspaceId *string `json:"workspace_id,omitempty"`

	// Playbook id.
	PlaybookId *string `json:"playbook_id,omitempty"`

	// dataclass id.
	DataclassId *string `json:"dataclass_id,omitempty"`

	// If the condition filter is enabled.
	RuleEnable *bool `json:"rule_enable,omitempty"`

	// If is enabled, false for disenabled, true for enabled
	Enabled *bool `json:"enabled,omitempty"`

	// 状态
	Status *string `json:"status,omitempty"`

	// 规则ID
	RuleId *string `json:"rule_id,omitempty"`

	// Strategy of action. event, timer
	TriggerType *string `json:"trigger_type,omitempty"`

	// if trigger when dataobject is created
	DataobjectCreate *bool `json:"dataobject_create,omitempty"`

	// if trigger when dataobject is updated
	DataobjectUpdate *bool `json:"dataobject_update,omitempty"`

	// if trigger when dataobject is deleted
	DataobjectDelete *bool `json:"dataobject_delete,omitempty"`

	// Strategy of action. sync or async
	ActionStrategy *string `json:"action_strategy,omitempty"`
}

func (o ModifyPlaybookVersionInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifyPlaybookVersionInfo struct{}"
	}

	return strings.Join([]string{"ModifyPlaybookVersionInfo", string(data)}, " ")
}
