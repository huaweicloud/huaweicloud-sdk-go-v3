package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ModifyPlaybookVersionInfo 剧本版本信息
type ModifyPlaybookVersionInfo struct {

	// 描述
	Description *string `json:"description,omitempty"`

	// 工作空间ID
	WorkspaceId *string `json:"workspace_id,omitempty"`

	// 剧本ID
	PlaybookId *string `json:"playbook_id,omitempty"`

	// 数据类ID
	DataclassId *string `json:"dataclass_id,omitempty"`

	// 是否启用触发条件过滤器
	RuleEnable *bool `json:"rule_enable,omitempty"`

	// 是否激活。(false:未激活, true:已激活)
	Enabled *bool `json:"enabled,omitempty"`

	// 状态(APPROVING:审核中，EDITING-编辑中，UNPASSED-审核未通过，PUBLISHED-已发布)
	Status *string `json:"status,omitempty"`

	// 规则ID
	RuleId *string `json:"rule_id,omitempty"`

	// 触发方式. EVENT--事件触发, TIMER--定时触发
	TriggerType *string `json:"trigger_type,omitempty"`

	// 数据对象是否创建时触发剧本
	DataobjectCreate *bool `json:"dataobject_create,omitempty"`

	// 数据对象是否更新时触发剧本
	DataobjectUpdate *bool `json:"dataobject_update,omitempty"`

	// 数据对象是否删除时触发剧本
	DataobjectDelete *bool `json:"dataobject_delete,omitempty"`

	// 执行策略. 目前仅支持异步并发执行，对应值为ASYNC
	ActionStrategy *string `json:"action_strategy,omitempty"`
}

func (o ModifyPlaybookVersionInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifyPlaybookVersionInfo struct{}"
	}

	return strings.Join([]string{"ModifyPlaybookVersionInfo", string(data)}, " ")
}
