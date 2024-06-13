package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RequirementsOverviewVo 实际的数据类型：单个对象，集合 或 NULL
type RequirementsOverviewVo struct {

	// 质量报告需求测试情况总数
	TotalNumber *int32 `json:"total_number,omitempty"`

	// 质量报告需求测试情况列表
	RequirementOverviewList *[]RequirementOverviewVo `json:"requirement_overview_list,omitempty"`
}

func (o RequirementsOverviewVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RequirementsOverviewVo struct{}"
	}

	return strings.Join([]string{"RequirementsOverviewVo", string(data)}, " ")
}
