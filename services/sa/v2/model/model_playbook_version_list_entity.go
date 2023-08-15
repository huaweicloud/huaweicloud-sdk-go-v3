package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PlaybookVersionListEntity Information of playbook version
type PlaybookVersionListEntity struct {

	// Id value
	Id *string `json:"id,omitempty"`

	// The description, display only
	Description *string `json:"description,omitempty"`

	// Create time
	CreateTime *string `json:"create_time,omitempty"`

	// Update time
	UpdateTime *string `json:"update_time,omitempty"`

	// Project id value
	ProjectId *string `json:"project_id,omitempty"`

	// Creator id value
	CreatorId *string `json:"creator_id,omitempty"`

	// Modifier id value
	ModifierId *string `json:"modifier_id,omitempty"`

	// Playbook id.
	PlaybookId *string `json:"playbook_id,omitempty"`

	// version
	Version *string `json:"version,omitempty"`

	// Run mode of this playbook. automatic, manual
	RunMode *string `json:"run_mode,omitempty"`

	// If is enabled, false for disenabled, true for enabled
	Enabled *bool `json:"enabled,omitempty"`

	// Status of approvement. editing, approving, unpassed, published
	Status *string `json:"status,omitempty"`

	// Strategy of action. sync or async
	ActionStrategy *string `json:"action_strategy,omitempty"`

	// If the condition filter is enabled.
	RuleEnable *bool `json:"rule_enable,omitempty"`

	// bind dataclass id
	DataclassId *string `json:"dataclass_id,omitempty"`

	// Strategy of action. event, timer
	TriggerType *string `json:"trigger_type,omitempty"`

	// if trigger when dataobject is created
	DataobjectCreate *bool `json:"dataobject_create,omitempty"`

	// if trigger when dataobject is updated
	DataobjectUpdate *bool `json:"dataobject_update,omitempty"`

	// if trigger when dataobject is deleted
	DataobjectDelete *bool `json:"dataobject_delete,omitempty"`

	// 版本类型
	VersionType *int32 `json:"version_type,omitempty"`

	// 过滤规则ID
	RuleId *string `json:"rule_id,omitempty"`

	// 数据类名称
	DataclassName *string `json:"dataclass_name,omitempty"`

	// 审核者
	ApproveName *string `json:"approve_name,omitempty"`

	// dataobject id
	DataobjectId *string `json:"dataobject_id,omitempty"`
}

func (o PlaybookVersionListEntity) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PlaybookVersionListEntity struct{}"
	}

	return strings.Join([]string{"PlaybookVersionListEntity", string(data)}, " ")
}
