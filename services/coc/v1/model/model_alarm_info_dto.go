package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AlarmInfoDto This is a alarm
type AlarmInfoDto struct {

	// 唯一标识ID
	Id *string `json:"id,omitempty"`

	// 名称
	Name *string `json:"name,omitempty"`

	// 严重类型，urgent（紧急）、major（重要）、minor（次要）、warn（提示）
	Importance *string `json:"importance,omitempty"`

	// 数据源
	ComeFrom *string `json:"come_from,omitempty"`

	// 数据源英文名
	ComeFromEn *string `json:"come_from_en,omitempty"`

	// 流转规则id
	TransferRule *string `json:"transfer_rule,omitempty"`

	// 流转规则名
	TransferRuleName *string `json:"transfer_rule_name,omitempty"`

	// 应用
	App *string `json:"app,omitempty"`

	// 状态，告警中alarming，已解决resolved
	Status *string `json:"status,omitempty"`

	// 责任人id
	Owner *string `json:"owner,omitempty"`

	// 责任人姓名
	OwnerName *string `json:"owner_name,omitempty"`

	// 责任人别名
	OwnerAlias *string `json:"owner_alias,omitempty"`

	// 收敛量
	ConvergeCount *int32 `json:"converge_count,omitempty"`

	// 关联事件id
	AssociateEventId *string `json:"associate_event_id,omitempty"`

	// 租户账号
	DomainId *string `json:"domain_id,omitempty"`

	// 创建人id
	Creator *string `json:"creator,omitempty"`

	// 创建人姓名
	CreatorName *string `json:"creator_name,omitempty"`

	// 创建时间
	CreateTime *int64 `json:"create_time,omitempty"`

	// 更新时间
	UpdateTime *int64 `json:"update_time,omitempty"`

	// 备注
	Remarks *string `json:"remarks,omitempty"`

	// 区域regionId
	Region *string `json:"region,omitempty"`

	// 任务类型 ：SCRIPT脚本，RUNBOOK作业
	TaskType *string `json:"task_type,omitempty"`

	// 任务类型分类 CUSTOMIZATION自定义，COMMUNAL公共
	AssociatedTaskType *string `json:"associated_task_type,omitempty"`

	// 脚本或作业id
	AssociatedTaskId *string `json:"associated_task_id,omitempty"`

	// 脚本或作业名称
	AssociatedTaskName *string `json:"associated_task_name,omitempty"`
}

func (o AlarmInfoDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AlarmInfoDto struct{}"
	}

	return strings.Join([]string{"AlarmInfoDto", string(data)}, " ")
}
