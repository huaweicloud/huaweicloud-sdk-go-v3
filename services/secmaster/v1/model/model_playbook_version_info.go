package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PlaybookVersionInfo 剧本版本详情信息
type PlaybookVersionInfo struct {

	// 剧本版本ID
	Id *string `json:"id,omitempty"`

	// 描述
	Description *string `json:"description,omitempty"`

	// 创建时间
	CreateTime *string `json:"create_time,omitempty"`

	// 更新时间
	UpdateTime *string `json:"update_time,omitempty"`

	// 项目ID
	ProjectId *string `json:"project_id,omitempty"`

	// 创建者ID
	CreatorId *string `json:"creator_id,omitempty"`

	// 修改者ID
	ModifierId *string `json:"modifier_id,omitempty"`

	// 剧本ID
	PlaybookId *string `json:"playbook_id,omitempty"`

	// 版本号
	Version *string `json:"version,omitempty"`

	// 是否启用。（true--已启用，false-未启用）
	Enabled *bool `json:"enabled,omitempty"`

	// 剧本版本状态，编辑中：EDITING 审核中：APPROVING 不通过：UNPASSED 已发布：PUBLISHED
	Status *string `json:"status,omitempty"`

	// 执行策略. 目前仅支持异步并发执行，对应值为ASYNC
	ActionStrategy *string `json:"action_strategy,omitempty"`

	// 剧本关联流程列表信息
	Actions *[]ActionInfo `json:"actions,omitempty"`

	// 是否启用触发条件过滤器
	RuleEnable *bool `json:"rule_enable,omitempty"`

	Rules *RuleInfo `json:"rules,omitempty"`

	// 数据类ID
	DataclassId *string `json:"dataclass_id,omitempty"`

	// 剧本触发方式(EVENT--事件触发，TIMER--定时触发)
	TriggerType *string `json:"trigger_type,omitempty"`

	// 标识数据对象是否创建时触发剧本
	DataobjectCreate *bool `json:"dataobject_create,omitempty"`

	// 标识数据对象是否更新时触发剧本
	DataobjectUpdate *bool `json:"dataobject_update,omitempty"`

	// 标识数据对象是否删除时触发剧本
	DataobjectDelete *bool `json:"dataobject_delete,omitempty"`

	// 版本类型(0--草稿版本，1--正式版本)
	VersionType *int32 `json:"version_type,omitempty"`

	// 过滤规则ID
	RuleId *string `json:"rule_id,omitempty"`

	// 数据类名称
	DataclassName *string `json:"dataclass_name,omitempty"`

	// 审核者
	ApproveName *string `json:"approve_name,omitempty"`
}

func (o PlaybookVersionInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PlaybookVersionInfo struct{}"
	}

	return strings.Join([]string{"PlaybookVersionInfo", string(data)}, " ")
}
