package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 测试计划详情
type TestPlanDetail struct {

	// 测试计划id
	PlanId *string `json:"plan_id,omitempty" xml:"plan_id"`

	// 测试计划名称
	Name *string `json:"name,omitempty" xml:"name"`

	// 测试计划开始时间
	StartDate *string `json:"start_date,omitempty" xml:"start_date"`

	// 测试计划截止时间
	EndDate *string `json:"end_date,omitempty" xml:"end_date"`

	// 测试计划实际完成时间（测试计划实际完成指测试计划下所有测试用例处于完成状态）
	FinishDate *string `json:"finish_date,omitempty" xml:"finish_date"`

	// 项目id
	ProjectId *string `json:"project_id,omitempty" xml:"project_id"`

	// 当前测试计划所处的阶段
	CurrentStage *string `json:"current_stage,omitempty" xml:"current_stage"`

	// 获取超期时间,正值表示已超期
	ExpireDay *string `json:"expire_day,omitempty" xml:"expire_day"`

	Creator *TestPlanDetailCreator `json:"creator,omitempty" xml:"creator"`

	Owner *TestPlanDetailOwner `json:"owner,omitempty" xml:"owner"`

	DesignStage *TestPlanDetailDesignStage `json:"design_stage,omitempty" xml:"design_stage"`

	ExecuteStage *TestPlanDetailExecuteStage `json:"execute_stage,omitempty" xml:"execute_stage"`

	ReportStage *TestPlanDetailReportStage `json:"report_stage,omitempty" xml:"report_stage"`

	Iteration *NameAndId `json:"iteration,omitempty" xml:"iteration"`
}

func (o TestPlanDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TestPlanDetail struct{}"
	}

	return strings.Join([]string{"TestPlanDetail", string(data)}, " ")
}
