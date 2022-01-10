package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 项目下某个测试计划操作历史列表
type TestPlanJournalList struct {
	// DevCloud项目id，项目唯一标识，固定长度32位字符

	ProjectId *string `json:"project_id,omitempty"`
	// 测试计划id

	PlanId *string `json:"plan_id,omitempty"`
	// 变更时间

	OperateTime *string `json:"operate_time,omitempty"`

	Operator *NameAndId `json:"operator,omitempty"`
	// 历史记录详情

	Detail *[]TestPlanJournalDetail `json:"detail,omitempty"`
}

func (o TestPlanJournalList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TestPlanJournalList struct{}"
	}

	return strings.Join([]string{"TestPlanJournalList", string(data)}, " ")
}
