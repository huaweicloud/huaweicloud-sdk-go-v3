package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TestReportInfoRequestBody 查询质量报告看板统计信息请求体
type TestReportInfoRequestBody struct {

	// 测试计划id,(plan_id和branch_id不能同时为空，优先取plan_id)
	PlanId *string `json:"plan_id,omitempty"`

	// 分支id,(plan_id和branch_id不能同时为空，优先取plan_id)
	BranchId *string `json:"branch_id,omitempty"`

	// 模块ID(查询未设置传入-2)
	ModuleId *string `json:"module_id,omitempty"`

	// 筛选迭代ID(查询未设置传入-2)
	FixedVersionId *string `json:"fixed_version_id,omitempty"`
}

func (o TestReportInfoRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TestReportInfoRequestBody struct{}"
	}

	return strings.Join([]string{"TestReportInfoRequestBody", string(data)}, " ")
}
