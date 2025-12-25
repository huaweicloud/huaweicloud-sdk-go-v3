package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreatePlaybookVersionInfo 创建剧本版本请求参数
type CreatePlaybookVersionInfo struct {

	// 描述
	Description *string `json:"description,omitempty"`

	// 工作空间ID
	WorkspaceId *string `json:"workspace_id,omitempty"`

	// 剧本ID
	PlaybookId *string `json:"playbook_id,omitempty"`

	// 关联流程列表
	Actions *[]CreateAction `json:"actions,omitempty"`

	// 数据类ID
	DataclassId string `json:"dataclass_id"`

	// 过滤规则是否启用
	RuleEnable *bool `json:"rule_enable,omitempty"`

	// 过滤规则ID
	RuleId *string `json:"rule_id,omitempty"`

	// 触发方式. EVENT--事件触发, TIMER--定时触发
	TriggerType *string `json:"trigger_type,omitempty"`

	// 标识数据对象是否创建时触发剧本
	DataobjectCreate *bool `json:"dataobject_create,omitempty"`

	// 标识数据对象是否更新时触发剧本
	DataobjectUpdate *bool `json:"dataobject_update,omitempty"`

	// 标识数据对象是否删除时触发剧本
	DataobjectDelete *bool `json:"dataobject_delete,omitempty"`

	// 执行策略. 目前仅支持异步并发执行，对应值为ASYNC
	ActionStrategy *string `json:"action_strategy,omitempty"`

	Rule *CreateRuleInfo `json:"rule,omitempty"`
}

func (o CreatePlaybookVersionInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePlaybookVersionInfo struct{}"
	}

	return strings.Join([]string{"CreatePlaybookVersionInfo", string(data)}, " ")
}
