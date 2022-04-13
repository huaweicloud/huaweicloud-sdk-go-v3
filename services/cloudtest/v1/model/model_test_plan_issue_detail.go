package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 测试计划需求列表详情
type TestPlanIssueDetail struct {
	// DevCloud项目id，项目唯一标识，固定长度32位字符

	ProjectId *string `json:"project_id,omitempty"`
	// 测试计划id

	PlanId *string `json:"plan_id,omitempty"`
	// 工作项id

	WorkitemId *string `json:"workitem_id,omitempty"`
	// 父工作项

	ParentIssue *string `json:"parent_issue,omitempty"`
	// 预计开始日期

	StartDate *string `json:"start_date,omitempty"`
	// 预计结束日期

	EndDate *string `json:"end_date,omitempty"`
	// 工作项名称

	Name *string `json:"name,omitempty"`
	// region信息

	RegionId *string `json:"region_id,omitempty"`

	Owner *NameAndId `json:"owner,omitempty"`

	Severity *NameAndId `json:"severity,omitempty"`

	Status *NameAndId `json:"status,omitempty"`

	Tracker *NameAndId `json:"tracker,omitempty"`

	Iteration *NameAndId `json:"iteration,omitempty"`

	Module *NameAndId `json:"module,omitempty"`
}

func (o TestPlanIssueDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TestPlanIssueDetail struct{}"
	}

	return strings.Join([]string{"TestPlanIssueDetail", string(data)}, " ")
}
