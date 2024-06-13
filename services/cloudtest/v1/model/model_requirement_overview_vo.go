package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RequirementOverviewVo 质量报告需求测试情况列表
type RequirementOverviewVo struct {

	// 需求名称
	Name *string `json:"name,omitempty"`

	// 需求id
	WorkitemId *string `json:"workitem_id,omitempty"`

	// 需求序列编号
	SequenceId *string `json:"sequence_id,omitempty"`

	// 看板需求id
	BoardId *string `json:"board_id,omitempty"`

	// 需求类型id
	TrackerId *string `json:"tracker_id,omitempty"`

	// 需求类型
	TrackerName *string `json:"tracker_name,omitempty"`

	// 需求关联用例总数
	RelateCaseNumber *int32 `json:"relate_case_number,omitempty"`

	CasePassVo *CasePassVo `json:"case_pass_vo,omitempty"`

	CaseExecuteVo *CaseExecuteVo `json:"case_execute_vo,omitempty"`

	// 需求关联缺陷总数
	RelateDefectNumber *int32 `json:"relate_defect_number,omitempty"`
}

func (o RequirementOverviewVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RequirementOverviewVo struct{}"
	}

	return strings.Join([]string{"RequirementOverviewVo", string(data)}, " ")
}
