package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 测试计划需求列表详情
type TestPlanIssueDetail struct {

	// DevCloud项目id，项目唯一标识，固定长度32位字符
	ProjectId *string `json:"project_id,omitempty" xml:"project_id"`

	// 测试计划id
	PlanId *string `json:"plan_id,omitempty" xml:"plan_id"`

	// 工作项id
	WorkitemId *string `json:"workitem_id,omitempty" xml:"workitem_id"`

	// 父工作项
	ParentIssue *string `json:"parent_issue,omitempty" xml:"parent_issue"`

	// 预计开始日期
	StartDate *string `json:"start_date,omitempty" xml:"start_date"`

	// 预计结束日期
	EndDate *string `json:"end_date,omitempty" xml:"end_date"`

	// 工作项名称
	Name *string `json:"name,omitempty" xml:"name"`

	// region信息
	RegionId *string `json:"region_id,omitempty" xml:"region_id"`

	Owner *NameAndId `json:"owner,omitempty" xml:"owner"`

	Severity *NameAndId `json:"severity,omitempty" xml:"severity"`

	Status *NameAndId `json:"status,omitempty" xml:"status"`

	Tracker *NameAndId `json:"tracker,omitempty" xml:"tracker"`

	Iteration *NameAndId `json:"iteration,omitempty" xml:"iteration"`

	Module *NameAndId `json:"module,omitempty" xml:"module"`
}

func (o TestPlanIssueDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TestPlanIssueDetail struct{}"
	}

	return strings.Join([]string{"TestPlanIssueDetail", string(data)}, " ")
}
