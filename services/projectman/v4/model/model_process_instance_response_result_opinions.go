package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ProcessInstanceResponseResultOpinions struct {

	// id
	Id *string `json:"id,omitempty"`

	// 类型
	Type *string `json:"type,omitempty"`

	// 状态
	State *string `json:"state,omitempty"`

	// 状态
	Status *string `json:"status,omitempty"`

	// 区域
	Region *string `json:"region,omitempty"`

	// 类型
	Category *string `json:"category,omitempty"`

	// 标题
	Title *string `json:"title,omitempty"`

	// 范围
	Rounds *string `json:"rounds,omitempty"`

	// 观察者
	Opinion *string `json:"opinion,omitempty"`

	// 描述
	Description *string `json:"description,omitempty"`

	// 修改人
	ModifiedBy *string `json:"modified_by,omitempty"`

	// 修改时间
	ModifiedDate *string `json:"modified_date,omitempty"`

	CreatedBy *ProcessInstanceResponseResultCreatedBy1 `json:"created_by,omitempty"`

	// 创建时间
	CreatedDate *string `json:"created_date,omitempty"`

	// 租户id
	TenantId *string `json:"tenant_id,omitempty"`

	// 项目ID
	DomainId *string `json:"domain_id,omitempty"`

	// 工作项类型
	IssueCategory *string `json:"issue_category,omitempty"`

	// 工作项ID
	IssueId *string `json:"issue_id,omitempty"`

	CurrOwner *ProcessInstanceResponseResultCurrOwner `json:"curr_owner,omitempty"`

	// 变更对象id
	CoId *string `json:"co_id,omitempty"`

	// 用户ID
	UserId *string `json:"user_id,omitempty"`

	// 评审工作项ID
	OpinionIssueId *string `json:"opinion_issue_id,omitempty"`

	// 评审工作项类型
	OpinionIssueCategory *string `json:"opinion_issue_category,omitempty"`
}

func (o ProcessInstanceResponseResultOpinions) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProcessInstanceResponseResultOpinions struct{}"
	}

	return strings.Join([]string{"ProcessInstanceResponseResultOpinions", string(data)}, " ")
}
