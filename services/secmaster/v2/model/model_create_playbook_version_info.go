package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Information of playbook version
type CreatePlaybookVersionInfo struct {

	// The description, display only
	Description *string `json:"description,omitempty"`

	// 工作空间id
	WorkspaceId *string `json:"workspace_id,omitempty"`

	// Playbook id.
	PlaybookId *string `json:"playbook_id,omitempty"`

	// Information of actions.
	Actions *[]ActionInfo `json:"actions,omitempty"`

	// dataclass id.
	DataclassId *string `json:"dataclass_id,omitempty"`

	// 过滤规则是否启用
	RuleEnable *bool `json:"rule_enable,omitempty"`

	// 过滤规则ID
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

func (o CreatePlaybookVersionInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePlaybookVersionInfo struct{}"
	}

	return strings.Join([]string{"CreatePlaybookVersionInfo", string(data)}, " ")
}
