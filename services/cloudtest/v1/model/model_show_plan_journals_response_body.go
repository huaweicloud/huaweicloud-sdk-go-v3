package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ShowPlanJournalsResponseBody struct {
	// DevCloud项目id，项目唯一标识，固定长度32位字符

	ProjectId *string `json:"project_id,omitempty"`
	// 测试计划id

	PlanId *string `json:"plan_id,omitempty"`
	// 变更时间

	OperateTime *string `json:"operate_time,omitempty"`

	Operator *NameAndId `json:"operator,omitempty"`
	// 历史记录详情

	Detail *[]ShowPlanJournalsResponseDetail `json:"detail,omitempty"`
}

func (o ShowPlanJournalsResponseBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPlanJournalsResponseBody struct{}"
	}

	return strings.Join([]string{"ShowPlanJournalsResponseBody", string(data)}, " ")
}
