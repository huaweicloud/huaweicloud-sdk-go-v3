package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PlaybookVersionListEntity 剧本版本详情信息
type PlaybookVersionListEntity struct {

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

	// 是否激活
	Enabled *bool `json:"enabled,omitempty"`

	// 状态. (EDITING--编辑中, APPROVING--审核中, UNPASSED--审核不通过, PUBLISHED--审核通过)
	Status *string `json:"status,omitempty"`

	// 执行策略. 目前仅支持异步并发执行，对应值为ASYNC
	ActionStrategy *string `json:"action_strategy,omitempty"`

	// 过滤规则是否启用
	RuleEnable *bool `json:"rule_enable,omitempty"`

	// 数据类ID
	DataclassId *string `json:"dataclass_id,omitempty"`

	// 触发方式. EVENT--事件触发, TIMER--定时触发
	TriggerType *string `json:"trigger_type,omitempty"`

	// 标识数据对象是否创建时触发剧本
	DataobjectCreate *bool `json:"dataobject_create,omitempty"`

	// 标识数据对象是否更新时触发剧本
	DataobjectUpdate *bool `json:"dataobject_update,omitempty"`

	// 标识数据对象是否删除时触发剧本
	DataobjectDelete *bool `json:"dataobject_delete,omitempty"`

	// 版本类型
	VersionType *int32 `json:"version_type,omitempty"`

	// 过滤规则ID
	RuleId *string `json:"rule_id,omitempty"`

	// 数据类名称
	DataclassName *string `json:"dataclass_name,omitempty"`

	// 审核者
	ApproveName *string `json:"approve_name,omitempty"`
}

func (o PlaybookVersionListEntity) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PlaybookVersionListEntity struct{}"
	}

	return strings.Join([]string{"PlaybookVersionListEntity", string(data)}, " ")
}
