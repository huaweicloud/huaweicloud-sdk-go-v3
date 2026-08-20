package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ShowIpdProcessInstancesResponseResultOpinions struct {

	// opinion主键。
	Id *string `json:"id,omitempty"`

	// 类型分类。
	Type *string `json:"type,omitempty"`

	// 数据状态。
	State *string `json:"state,omitempty"`

	// opinion状态。
	Status *string `json:"status,omitempty"`

	// 区域。
	Region *string `json:"region,omitempty"`

	// 类型。
	Category *string `json:"category,omitempty"`

	// 标题。
	Title *string `json:"title,omitempty"`

	// 评审轮次。
	Rounds *string `json:"rounds,omitempty"`

	// 评审意见。
	Opinion *string `json:"opinion,omitempty"`

	// 描述。
	Description *string `json:"description,omitempty"`

	// 修改人。
	ModifiedBy *string `json:"modified_by,omitempty"`

	// 修改时间。
	ModifiedDate *string `json:"modified_date,omitempty"`

	CreatedBy *UserObject `json:"created_by,omitempty"`

	// 创建时间。
	CreatedDate *string `json:"created_date,omitempty"`

	// 租户ID。
	TenantId *string `json:"tenant_id,omitempty"`

	// 项目空间ID。
	DomainId *string `json:"domain_id,omitempty"`

	// 对象类型。
	IssueCategory *string `json:"issue_category,omitempty"`

	// 对象ID。
	IssueId *string `json:"issue_id,omitempty"`

	CurrOwner *UserObject `json:"curr_owner,omitempty"`

	// 变更对象ID。
	CoId *string `json:"co_id,omitempty"`

	// 用户ID。
	UserId *string `json:"user_id,omitempty"`

	// 评审工作项ID。
	OpinionIssueId *string `json:"opinion_issue_id,omitempty"`

	// 评审工作项类型。
	OpinionIssueCategory *string `json:"opinion_issue_category,omitempty"`
}

func (o ShowIpdProcessInstancesResponseResultOpinions) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowIpdProcessInstancesResponseResultOpinions struct{}"
	}

	return strings.Join([]string{"ShowIpdProcessInstancesResponseResultOpinions", string(data)}, " ")
}
