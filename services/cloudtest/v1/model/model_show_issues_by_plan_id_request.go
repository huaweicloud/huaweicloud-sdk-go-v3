package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowIssuesByPlanIdRequest struct {

	// 测试计划id，测试计划的唯一标识，长度11-34位字符
	PlanId string `json:"plan_id"`

	// 起始偏移量，表示从此偏移量开始查询， offset大于等于0
	Offset int64 `json:"offset"`

	// 每页显示的条目数量,最大支持200条
	Limit int64 `json:"limit"`
}

func (o ShowIssuesByPlanIdRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowIssuesByPlanIdRequest struct{}"
	}

	return strings.Join([]string{"ShowIssuesByPlanIdRequest", string(data)}, " ")
}
