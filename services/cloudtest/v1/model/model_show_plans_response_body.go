package model

import (
	"encoding/json"

	"strings"
)

type ShowPlansResponseBody struct {
	// 测试计划id

	PlanId *string `json:"plan_id,omitempty"`
	// 测试计划名称

	Name *string `json:"name,omitempty"`
	// 测试计划开始日期

	StartDate *string `json:"start_date,omitempty"`
	// 测试计划截止日期

	EndDate *string `json:"end_date,omitempty"`
	// 测试计划实际完成日期（测试计划实际完成指测试计划下所有测试用例处于完成状态）

	FinishDate *string `json:"finish_date,omitempty"`
	// 项目id

	ProjectId *string `json:"project_id,omitempty"`
	// 当前测试计划所处的阶段

	CurrentStage *string `json:"current_stage,omitempty"`
	// 获取超期时间,正值表示已超期

	ExpireDay *string `json:"expire_day,omitempty"`

	Owner *ShowPlansResponseOwner `json:"owner,omitempty"`

	DesignStage *ShowPlansResponseDesignStage `json:"design_stage,omitempty"`

	ExecuteStage *ShowPlansResponseExecuteStage `json:"execute_stage,omitempty"`

	ReportStage *ShowPlansResponseReportStage `json:"report_stage,omitempty"`
}

func (o ShowPlansResponseBody) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowPlansResponseBody struct{}"
	}

	return strings.Join([]string{"ShowPlansResponseBody", string(data)}, " ")
}
